package main

import (
	"fmt"
)

func main() {
	array1 := [3]int{1, 2, 4}
	var slice1 []int

	//array1 = append(array1, 3, 5, 6)
	slice1 = append(slice1, 4, 56, 3)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)

	// secao 5 aula 37
}
