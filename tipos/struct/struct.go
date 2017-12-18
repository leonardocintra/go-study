package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// metodo: funcao com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.78,
		desconto: 0.05,
	}

	// forma reduzida
	produto2 := produto{nome: "Borracha", preco: 2, desconto: 0.09}

	fmt.Println(produto1, produto1.precoComDesconto())
	fmt.Println(produto2, produto2.precoComDesconto())
}
