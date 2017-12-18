package main

import (
	"fmt"
)

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1000, quantidade: 2, preco: 12.50},
			item{1020, 3, 23.56},
			item{2222, 3, 23.56},
			item{3333, 2, 23.56},
			item{4444, 7, 23.56},
		},
	}

	fmt.Printf("Valor total do pedido: R$ %.2f", pedido.valorTotal())
}
