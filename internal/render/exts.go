package render

import (
	"path/filepath"
	"strings"
)

func Exts(name string) []string {
	var exts []string

	ext := filepath.Ext(name)

	for ext != "" {
		exts = append([]string{ext}, exts...)
		name = strings.TrimSuffix(name, ext)
		ext = filepath.Ext(name)
	}
	return exts
}

// HasExt checks if a file has ANY of the
// extensions passed in. If no extensions
// are given then `true` is returned
func HasExt(name string, ext ...string) bool {
	if len(ext) == 0 || ext == nil {
		return true
	}
	for _, xt := range ext {
		xt = strings.TrimSpace(xt)
		if xt == "*" || xt == "*.*" {
			return true
		}
		for _, x := range Exts(name) {
			if x == xt {
				return true
			}
		}
	}
	return false
}

// StripExt from a File and return a new one
func StripExt(name, ext string) string {
	name = strings.Replace(name, ext, "", -1)
	return name
}
