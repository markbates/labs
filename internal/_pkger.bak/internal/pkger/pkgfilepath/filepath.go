package pkgfilepath

import (
	"os"
	"path/filepath"

	"github.com/markbates/labs/internal/pkger/internal/pkger/pkgos"
)

func Walk(root string, wf filepath.WalkFunc) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		info = pkgos.NewFileInfo(info)
		return wf(path, info, err)
	})

	return err
}
