package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iterecao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Leonardo"," Porque voce nao fala comigo?", 3)
	// fale("João", "So posso falar depois de voce", 1)

	// go fale("Leonardo", " Porque voce nao fala comigo?", 3)
	// go fale("João", "So posso falar depois de voce", 1)

	go fale("Leonardo", " Porque voce nao fala comigo?", 10)
	fale("João", "So posso falar depois de voce", 6)

	fmt.Println("FIM")
}
