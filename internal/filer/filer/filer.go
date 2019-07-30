package filer

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/markbates/labs/internal/filer/filer/internal/filex"
)

type Filer struct {
	Args     []string
	IO       IO
	WalkFunc filepath.WalkFunc
}

func (f Filer) Run() error {
	root, err := modRoot()
	if err != nil {
		return err
	}

	args := make([]string, len(f.Args))
	copy(args, f.Args)
	if len(args) > 0 {
		root = args[0]
		args = args[1:]
	}

	wf := f.WalkFunc
	if wf == nil {
		wf = func(path string, info os.FileInfo, err error) error {
			return err
		}
	}

	wf = filex.SkipDir(".git", wf)
	wf = filex.SkipDir("testdata", wf)
	wf = filex.SkipDir("node_modules", wf)
	wf = filex.SkipFilePrefix("_", wf)

	if err := filepath.Walk(root, defaultWalk(wf)); err != nil {
		return err
	}
	return nil
}

func modRoot() (string, error) {
	c := exec.Command("go", "env", "GOMOD")
	b, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}

	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return "", fmt.Errorf("could not file GOMOD")
	}

	return filepath.Dir(string(b)), nil
}

func defaultWalk(wf filepath.WalkFunc) filepath.WalkFunc {
	mp := newFileMap()
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if err := wf(path, info, err); err != nil {
			return err
		}

		fi := NewFileInfo(info)
		f := &File{
			info: fi,
		}

		fmt.Println(path, f)
		mp.Store(path, f)

		return nil
	}
}

// return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
//
// 	if !info.IsDir() {
// 		return nil
// 	}
// 	oncer.Do(path, func() {
// 		fmt.Println(path)
// 		cfg := &packages.Config{TheMode: packages.NeedFiles | packages.NeedSyntax}
// 		var pkgs []*packages.Package
// 		args = append([]string{path}, args...)
// 		pkgs, err = packages.Load(cfg, args...)
// 		if err != nil {
// 			return
// 		}
// 		if packages.PrintErrors(pkgs) > 0 {
// 			return
// 		}
//
// 		// Print the names of the source files
// 		// for each package listed on the command line.
// 		for _, pkg := range pkgs {
// 			// fmt.Println(pkg.ID, pkg.GoFiles)
// 			fmt.Println(pkg.ID)
// 		}
// 	})
// 	return err
// })
