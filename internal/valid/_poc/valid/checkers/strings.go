package checkers

import (
	"context"
	"fmt"
	"validator/valid"

	"github.com/google/go-cmp/cmp"
)

func StringsEqual(c valid.Checker, key string, a, b string) valid.Checker {
	return WithFunc(c, func(ctx context.Context) error {
		x := cmp.Equal(a, b)
		if !x {
			c.Add(key, fmt.Sprintf("%q does not equal %q", a, b))
		}
		return c.Validate(ctx)
	})
	return c
}
