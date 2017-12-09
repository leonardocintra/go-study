package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(100)
}

func main() {
	if i := numeroAleatorio(); i > 50 { // tbm suportado no switch
		fmt.Println("Ganho!")
	} else {
		fmt.Println("Perdeu troxa")
	}
}
