package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 3, 5} // array
	s1 := []int{4, 5, 6}  // slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{5, 4, 3, 2, 1}

	// SLICE NÂO EH ARRAY! Slice define um pedeço do array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	// novo slice mas  para o mesmo array
	s3 := a2[:3]
	fmt.Println(a2, s3)

	/*
		slice pode ser imaginad como:
			tamanho e um ponteiro para o elemento de um array
			tambem pode fazer um slice de um slice
	*/

}
