package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros interios
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	// sem sinal (so positivos) uint8 utin16 utin32 utin64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	il := math.MaxInt64
	fmt.Println("O valor maximo de int64 é", il)
	fmt.Println("O tipo de il é", reflect.TypeOf(il))

	// string com multimplas linhas usar `crase`
}
