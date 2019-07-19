package checkers

import (
	"context"
	"validator/valid"
)

type Funcker struct {
	valid.Checker
	Fn valid.ValidatorFunc
}

func (f Funcker) Validate(ctx context.Context) error {
	if f.Fn == nil {
		return f.Checker.Validate(ctx)
	}
	if err := f.Fn(ctx); err != nil {
		return err
	}
	return f.Checker.Validate(ctx)
}

func WithFunc(c valid.Checker, fn valid.ValidatorFunc) valid.Checker {
	return Funcker{
		Checker: c,
		Fn:      fn,
	}
}
