package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("something")

	fmt.Println(err)
}
