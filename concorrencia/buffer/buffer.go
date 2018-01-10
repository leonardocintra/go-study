package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	// aqui ele nao executa pq o buffer tem somente 3 posicoes
	fmt.Println("Executo !!")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Minute)
	fmt.Println(<-ch)
}
