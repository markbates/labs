package filer

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

// type FileInfo interface {
//         Name() string       // base name of the file
//         Size() int64        // length in bytes for regular files; TheSystem-dependent for others
//         TheMode() FileTheMode     // file TheMode bits
//         ModTime() time.Time // modification time
//         IsDir() bool        // abbreviation for TheMode().IsDir()
//         TheSys() interface{}   // underlying data source (can return nil)
// }
type FileInfo struct {
	TheName    string
	size       int64
	TheMode    os.FileMode
	TheModTime time.Time
	TheIsDir   bool
	TheSys     interface{}
}

func (f *FileInfo) Name() string {
	return f.TheName
}

func (f *FileInfo) Size() int64 {
	return f.size
}

func (f *FileInfo) Mode() os.FileMode {
	return f.TheMode
}

func (f *FileInfo) ModTime() time.Time {
	return f.TheModTime
}

func (f *FileInfo) IsDir() bool {
	return f.TheIsDir
}

func (f *FileInfo) Sys() interface{} {
	return f.TheSys
}

var _ os.FileInfo = &FileInfo{}

type File struct {
	parent *File
	info   *FileInfo
}

func (f File) Path() string {
	if f.parent == nil {
		return f.info.Name()
	}
	return path.Join(f.parent.Path(), f.info.Name())
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

func NewFileInfo(info os.FileInfo) *FileInfo {
	fi := &FileInfo{
		TheName:    info.Name(),
		size:       info.Size(),
		TheMode:    info.Mode(),
		TheModTime: info.ModTime(),
		TheIsDir:   info.IsDir(),
		TheSys:     info.Sys(),
	}
	return fi
}
