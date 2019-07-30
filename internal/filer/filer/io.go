package filer

import (
	"io"
	"os"
)

type IO interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
}

type FuncIO struct {
	InFn  func() io.Reader
	OutFn func() io.Writer
	ErrFn func() io.Writer
}

func (f FuncIO) In() io.Reader {
	if f.InFn == nil {
		return os.Stdin
	}
	return f.InFn()
}

func (f FuncIO) Out() io.Writer {
	if f.OutFn == nil {
		return os.Stdout
	}
	return f.OutFn()
}

func (f FuncIO) Err() io.Writer {
	if f.ErrFn == nil {
		return os.Stderr
	}
	return f.ErrFn()
}
