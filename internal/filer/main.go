package main

import (
	"flag"
	"log"

	"github.com/markbates/labs/internal/filer/filer"
)

func main() {
	flag.Parse()
	f := filer.Filer{
		Args: flag.Args(),
		IO:   filer.FuncIO{},
	}
	err := f.Run()
	if err != nil {
		log.Fatal(err)
	}
}
