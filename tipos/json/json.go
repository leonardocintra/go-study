package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	fmt.Println("Ronaldo!")

	// struc para json
	p1 := produto{
		1,
		"Notebook",
		1999.98,
		[]string{"Promocao", "Eletronico"},
	}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"ID":12,"nome":"caneta","preco":9.98,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])

}
