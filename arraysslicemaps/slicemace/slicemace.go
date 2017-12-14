package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 11, 22, 33, 44, 55, 66, 77, 88, 65, 32)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 97)
	s = append(s, 96)
	s = append(s, 95)
	fmt.Println(s, len(s), cap(s))
}
