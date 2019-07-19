package models

import (
	"context"
	"validator/valid"
	"validator/valid/checkers"
)

type Widget struct {
	Name string
	Age  int
}

func (w *Widget) BeforeSave(ctx context.Context) error {
	c := valid.Background()

	// c = checkers.StringsEqual(c, "name", w.Name, "Mark")
	c = checkers.IntsEqual(c, "age", w.Age, 42)

	c = checkers.WithFunc(c, func(cx context.Context) error {
		// Using Pop:
		// tx, ok := cx.Value("tx").(*pop.Connection)
		// if !ok {
		// 	return fmt.Errorf("no transaction found")
		// }
		return nil
	})

	return c.Validate(ctx)
}
