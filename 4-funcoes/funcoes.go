package main

import (
	"fmt"
)

func main() {
	fmt.Println(somar(10, 20))

	var f = func(n3 int8, n4 int8) {
		fmt.Print("o Resultado é: ")
		fmt.Print(somar(n3, n4))
	}

	f(40, 20)
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
