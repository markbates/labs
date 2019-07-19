package valid

import (
	"context"
)

type Validator interface {
	Validate(context.Context) error
}

type ValidatorFunc func(context.Context) error

func (v ValidatorFunc) Validate(ctx context.Context) error {
	return v(ctx)
}
