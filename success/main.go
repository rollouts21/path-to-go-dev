package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write command")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода")
	}

	text := scanner.Text()

	fileds := strings.Fields(text)
	fmt.Println(fileds)
}
