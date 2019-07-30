package main

import (
	"context"
	"fmt"
	"log"
	"validator/valid"
	"validator/valid/checkers"
	"validator/valid/models"
)

type Beatle struct {
	Name       string
	Instrument string
}

func main() {
	ctx := context.Background()
	c := valid.Background()

	c = checkers.StringsEqual(c, "letters", "b", "A")
	c = checkers.StringsEqual(c, "letters", "A", "A")

	if err := c.Validate(ctx); err != nil {
		fmt.Println(err)
	}

	w := &models.Widget{
		Name: "Janis",
	}

	err := w.BeforeSave(ctx)
	if err == nil {
		return
	}

	ve, ok := err.(*valid.Errors)
	if !ok {
		log.Fatal(err)
	}

	ve.Range(func(key string, values []string) bool {
		fmt.Printf("%s:\n", key)
		for _, s := range values {
			fmt.Printf("  * %q\n", s)
		}
		return true
	})

}
