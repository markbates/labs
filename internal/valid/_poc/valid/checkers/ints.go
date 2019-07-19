package checkers

import (
	"context"
	"fmt"
	"validator/valid"
)

func IntsEqual(c valid.Checker, key string, a, b int) valid.Checker {
	return WithFunc(c, func(ctx context.Context) error {
		if a != b {
			c.Add(key, fmt.Sprintf("%d does not equal %d", a, b))
		}
		return c.Validate(ctx)
	})
	return c
}
