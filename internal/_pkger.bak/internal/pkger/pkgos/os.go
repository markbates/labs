package pkgos

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

type pkgOS struct {
	modRoot string
	err     error
}

var oreos = func() *pkgOS {
	p := &pkgOS{}
	mr, err := modRoot()
	if err != nil {
		p.err = err
		return p
	}
	p.modRoot = mr
	return p
}()

func modRoot() (string, error) {
	c := exec.Command("go", "env", "GOMOD")
	b, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}

	b = bytes.TrimSpace(b)
	if len(b) == 0 {
		return "", fmt.Errorf("the `go env GOMOD` was empty/modules are required")
	}

	return filepath.Dir(string(b)), nil
}

func Getwd() (string, error) {
	if oreos.err != nil {
		return "", oreos.err
	}
	return oreos.modRoot, nil
}
