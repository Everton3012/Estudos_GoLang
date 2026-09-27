package main

import "fmt"

func alunoAprovado(nota float64) bool {
	if nota >= 7 {
		return true
	}
	return false
}

func main() {
	fmt.Println("defer em Go")

	defer fmt.Println(alunoAprovado(5))

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("Fim do programa")
}
