package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println("Valor de i", i)
		i++
	}

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print("Par")
		} else {
			fmt.Print("Impar")
		}
	}

	for {
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second * 5)
	}
}
