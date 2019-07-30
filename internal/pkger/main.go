package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/here"
)

const def = "github.com/gobuffalo/buffalo:server.go"

func main() {
	args := os.Args[1:]
	p := def
	if len(args) > 0 {
		p = args[0]
	}
	f, err := Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if info.IsDir() {
		fmt.Println(info.Name())
		return
	}
	io.Copy(os.Stdout, f)
}

func Open(p string) (*os.File, error) {
	res := strings.Split(p, ":")

	if len(res) < 1 {
		return nil, fmt.Errorf("could not parse %q (%d)", res, len(res))
	}

	var pt Path
	if len(res) == 1 {
		if strings.HasPrefix(res[0], "/") {
			pt.Name = res[0]
		} else {
			pt.Pkg = res[0]
		}
	} else {
		pt.Pkg = res[0]
		pt.Name = res[1]
	}

	// fmt.Println("pt", pt)

	info, err := here.Package(pt.Pkg)
	if err != nil {
		return nil, err
	}
	// fmt.Println(info)

	if len(pt.Name) == 0 {
		return os.Open(info.Dir)
	}

	return os.Open(filepath.Join(info.Dir, pt.Name))

	// cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedName}
	// pkgs, err := packages.Load(cfg, pt.Pkg)
	//
	// if err != nil {
	// 	return nil, err
	// }
	//
	// for _, pkg := range pkgs {
	// 	// fmt.Println(pkg.ID, pkg.OtherFiles)
	// 	// fmt.Println(pkg.ID, pkg.GoFiles)
	// }

	// return b, nil
}

type Path struct {
	Pkg  string
	Name string
}

// type File struct{}
