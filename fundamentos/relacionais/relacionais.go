package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Banana" == "Banana")

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))
}
