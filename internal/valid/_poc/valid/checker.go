package valid

import "context"

type Checker interface {
	Validator
	Errors() Errors
	Add(string, string)
	Set(string, []string)
}

func Background() Checker {
	return &backCheck{}
}

type backCheck struct {
	errors *Errors
}

func (bc backCheck) Errors() Errors {
	if bc.errors == nil {
		bc.errors = &Errors{}
	}
	return *bc.errors
}

func (bc *backCheck) Add(key string, value string) {
	if bc.errors == nil {
		bc.errors = &Errors{}
	}
	bc.errors.Add(key, value)
}

func (bc *backCheck) Set(key string, values []string) {
	if bc.errors == nil {
		bc.errors = &Errors{}
	}
	bc.errors.Set(key, values)
}

func (bc backCheck) Validate(ctx context.Context) error {
	if bc.errors == nil {
		bc.errors = &Errors{}
	}
	if len(bc.errors.errors) > 0 {
		return bc.errors
	}
	return nil
}
