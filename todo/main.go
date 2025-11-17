package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := SecondFunc(1); err != nil {
		var wErr SomethinError
		if errors.As(err, &wErr) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}

func FirstFunc(n int) error {
	if n == 1 {
		return ErrFirst
	}

	return SomethinError{n}
}

func SecondFunc(n int) error {
	if err := FirstFunc(n); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
