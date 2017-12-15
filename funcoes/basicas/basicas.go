package main

import (
	"fmt"
)

func f1() {
	fmt.Println("F1")
}

func f3(p1, p2 string) string {
	return fmt.Sprintf("F4 %s %s", p1, p2)
}

// funcao com retornos multiplos
func f5() (string, string) {
	return "Leonard", "Juliana"
}

func main() {
	f1()
	f3("ronaldo", "gaucho")
	r1, r2 := f5()
	fmt.Println(r1, r2)
}
