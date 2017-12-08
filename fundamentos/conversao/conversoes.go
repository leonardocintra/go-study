package main

import (
	"fmt"
	c "strconv"
)

func main() {
	x := 2.6
	fmt.Print(x)

	// cuidado
	fmt.Println("Teste " + string(123))

	// para converter para string tem que fazer assim
	fmt.Println("Teste " + c.Itoa(123))

	// string para int
	num, _ := c.Atoi("123")
	fmt.Println(num - 121)

	b, _ := c.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
