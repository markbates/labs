package render

import (
	"io"
	"strings"
)

type Data map[string]interface{}

type RenderFunc func(string, io.Writer, Data) error

func Plush(r RenderFunc) RenderFunc {
	return func(name string, w io.Writer, data Data) error {
		if strings.Contains(name, ".plush") {
			// do some work
		}
		return r(name, w, data)
	}
}
