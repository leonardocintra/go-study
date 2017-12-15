package main

import (
	"fmt"
)

var soma = func(a, b int) int {
	return b + a
}

func main() {
	fmt.Println(soma(2, 4))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(32, 39))
}
