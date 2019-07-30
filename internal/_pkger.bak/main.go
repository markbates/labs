package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/markbates/labs/internal/pkger/internal/pkger/pkgfilepath"
	"github.com/markbates/labs/internal/pkger/internal/pkger/pkgos"
)

var skip = regexp.MustCompile(`\A(\_|\.)`)

func main() {
	rd := flag.Bool("read", false, "read .index.json")
	flag.Parse()
	args := flag.Args()
	if *rd {
		readIndex(args)
		return
	}
	writeIndex(args)
}

func readIndex(args []string) {
	root, err := pkgos.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("root", root)

	f, err := os.Open(".index.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	index := map[string]*pkgos.File{}

	if err := json.NewDecoder(f).Decode(&index); err != nil {
		log.Fatal("&index ", err)
	}

	if len(args) > 0 {
		for _, a := range args {
			pf := index[a]
			if pf == nil {
				panic(a)
			}
			b, err := ioutil.ReadAll(pf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}
		return
	}
	for k, f := range index {
		fmt.Println(k, f)
	}
}

func writeIndex(args []string) {
	root, err := pkgos.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("root", root)

	files := map[string]*pkgos.File{}

	err = pkgfilepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)

		if skip.MatchString(base) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		f, err := pkgos.Open(path)
		if err != nil {
			return err
		}

		// defer f.Close()
		key := strings.TrimPrefix(path, root)
		fmt.Println("key: ", key)
		files[key] = f

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(filepath.Join(root, ".index.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(files)
	if err != nil {
		log.Fatal(err)
	}
}
