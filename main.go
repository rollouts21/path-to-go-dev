package main

import "fmt"

func main() {
	number := 10

	pointer := &number

	fmt.Println(pointer)

	foo(pointer)
}

func foo(n *int) {
	fmt.Println(n)
	fmt.Println(*n)
}
