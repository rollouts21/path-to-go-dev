package main

import "fmt"

func main() {
	a := 5.0
	b := 3.0

	fmt.Println(sum(a, b))
	fmt.Println(minus(a, b))
	fmt.Println(ym(a, b))
	if b > 0 {
		fmt.Println(del(a, b))
	} else {
		fmt.Println("На ноль делить нельзя!!!!")
	}
}

func sum(a float64, b float64) float64 {
	return a + b
}

func minus(a float64, b float64) float64 {
	return a - b
}

func ym(a float64, b float64) float64 {
	return a * b
}

func del(a float64, b float64) float64 {
	return a / b
}
