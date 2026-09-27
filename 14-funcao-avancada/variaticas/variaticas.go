package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções Variáticas em Go")
	fmt.Println("Soma:", soma())
	fmt.Println("Soma:", soma(1, 2))
	fmt.Println("Soma:", soma(1, 2, 3))
	fmt.Println("Soma:", soma(1, 2, 3, 4, 5))
	fmt.Println("Soma:", soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println("Soma:", soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15))
	fmt.Println("Soma:", soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))

	escrever("Número:", 1, 2, 3, 4, 5)
	escrever("Número:", 6, 7, 8, 9, 10)
}
