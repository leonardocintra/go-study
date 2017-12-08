package main

import "fmt"

func main() {
	fmt.Print("Mesmo")
	fmt.Print(" linha")

	fmt.Printf("Nova Linha")

	x := 32.32432

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 2.9894
	c := false
	d := "ronaldo"
	fmt.Printf("\n %d %f, %t, %s", a, b, c, d)

}

