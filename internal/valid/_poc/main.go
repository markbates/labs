package main

import (
	"context"
	"fmt"
	"log"
	"validator/valid"
	"validator/valid/checkers"
	"validator/valid/models"
)

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

	if err := w.BeforeSave(ctx); err != nil {
		ve, ok := err.(*valid.Errors)
		if !ok {
			log.Fatal(err)
		}
		ve.Range(func(key string, values []string) bool {
			fmt.Println(key, values)
			return true
		})
	}

}
