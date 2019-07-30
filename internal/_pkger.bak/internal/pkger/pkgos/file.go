package pkgos

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const timeFmt = time.RFC3339Nano

type FileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func (f *FileInfo) String() string {
	b, _ := json.MarshalIndent(f, "", "  ")
	return string(b)
}

func (f *FileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":    f.name,
		"size":    f.size,
		"mode":    f.mode,
		"modTime": f.modTime.Format(timeFmt),
		"isDir":   f.isDir,
		"sys":     f.sys,
	})
}

func (f *FileInfo) UnmarshalJSON(b []byte) error {
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	var ok bool

	f.name, ok = m["name"].(string)
	if !ok {
		return fmt.Errorf("could not determine name %q", m["name"])
	}

	size, ok := m["size"].(float64)
	if !ok {
		return fmt.Errorf("could not determine size %q", m["size"])
	}
	f.size = int64(size)

	mode, ok := m["mode"].(float64)
	if !ok {
		return fmt.Errorf("could not determine mode %q", m["mode"])
	}
	f.mode = os.FileMode(mode)

	modTime, ok := m["modTime"].(string)
	if !ok {
		return fmt.Errorf("could not determine modTime %q", m["modTime"])
	}
	t, err := time.Parse(timeFmt, modTime)
	if err != nil {
		return err
	}
	f.modTime = t

	f.isDir, ok = m["isDir"].(bool)
	if !ok {
		return fmt.Errorf("could not determine isDir %q", m["isDir"])
	}
	f.sys = m["sys"]
	return nil
}

func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) Size() int64 {
	return f.size
}

func (f *FileInfo) Mode() os.FileMode {
	return f.mode
}

func (f *FileInfo) ModTime() time.Time {
	return f.modTime
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Sys() interface{} {
	return f.sys
}

var _ os.FileInfo = &FileInfo{}

func Open(name string) (*File, error) {
	osf, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer osf.Close()

	info, err := osf.Stat()
	if err != nil {
		return nil, err
	}

	f := &File{
		info: NewFileInfo(info),
	}

	if !info.IsDir() {
		b, err := ioutil.ReadAll(osf)
		if err != nil {
			return nil, err
		}
		f.original = b
	}
	return f, nil
}

type File struct {
	info     *FileInfo
	original []byte
	index    int
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
	f.index = int(offset)
	return offset, nil
}

func (f File) Name() string {
	return f.info.Name()
}

func (f File) String() string {
	if f.info == nil {
		return ""
	}
	b, _ := json.MarshalIndent(f.info, "", "  ")
	return string(b)
}

func (f *File) Read(p []byte) (int, error) {
	// fmt.Println("read", f.Name())

	ln := len(f.original)
	pi := len(p)
	ei := f.index + pi

	// fmt.Printf("name %s | len %d | pi %d | index %d | [%d:%d] |\n", f.Name(), ln, pi, f.index, f.index, ei)

	if f.index >= ln {
		return 0, io.EOF
	}

	if ln <= pi {
		n := copy(p, f.original)
		f.index = ln
		return n, io.EOF
	}

	n := copy(p, f.original[f.index:ei])
	f.index = ei

	return n, nil

	// ln := len(f.original)
	// fmt.Println(f.Name(), ln, len(p))

	// if ln < len(p) {
	// 	return copy(p, f.original), io.EOF
	// }

	// var r io.Reader = bytes.NewReader(f.original)
	// var _ = r
	// bb := bytes.NewBuffer(p)
	// i, err := io.Copy(bb, r)
	// if i == 0 || err == io.EOF || int(i) == len(f.original) {
	// 	return int(i), io.EOF
	// }
	// return int(i), err

	// fmt.Println(len(f.original), len(p))
	// if len(f.original) == 0 {
	// 	f, err := os.Open(f.Name())
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	defer f.Close()
	// 	r = f
	// }
	//
	// bb := bytes.NewBuffer(p)

	// bb := bytes.NewBuffer(p)
	// i, err := io.Copy(bb, r)
	// if i == 0 || err == io.EOF || int(i) == len(f.original) {
	// 	return int(i), io.EOF
	// }
	// return int(i), err
}

func (f File) MarshalJSON() ([]byte, error) {
	var data []byte
	m := map[string]interface{}{
		"info": f.info,
		"data": data,
	}

	if f.info.IsDir() {
		return json.Marshal(m)
	}

	data, err := ioutil.ReadAll(&f)
	if err != nil {
		return nil, err
	}
	m["data"] = data
	return json.Marshal(m)

}

func (f *File) UnmarshalJSON(b []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	rm, ok := m["info"]
	if !ok {
		return fmt.Errorf("missing FileInfo")
	}

	f.info = &FileInfo{}
	if err := json.Unmarshal(rm, f.info); err != nil {
		return err
	}

	rm, ok = m["data"]
	if !ok {
		return fmt.Errorf("missing data")
	}

	f.original = rm
	return nil
}

func NewFileInfo(info os.FileInfo) *FileInfo {
	fi := &FileInfo{
		name:    info.Name(),
		size:    info.Size(),
		mode:    info.Mode(),
		modTime: info.ModTime(),
		isDir:   info.IsDir(),
		sys:     info.Sys(),
	}
	return fi
}
