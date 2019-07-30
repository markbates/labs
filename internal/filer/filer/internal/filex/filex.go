package filex

import (
	"os"
	"path/filepath"
	"strings"
)

func SkipDir(name string, fn filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		if base == name {
			return filepath.SkipDir
		}
		return fn(path, info, err)
	}
}

func SkipBase(name string, fn filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		if base == name {
			return nil
		}
		return fn(path, info, err)
	}
}

func SkipSuffix(name string, fn filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, name) {
			return nil
		}

		return fn(path, info, err)
	}
}

func SkipFilePrefix(name string, fn filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(filepath.Base(path), name) {
			return nil
		}

		return fn(path, info, err)
	}
}
