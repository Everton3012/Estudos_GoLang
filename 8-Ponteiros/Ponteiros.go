package main

import "fmt"

func main() {
	fmt.Println("Ponteiros em Go")

	var numero int = 10
	var ponteiro *int = &numero

	fmt.Println("Valor de numero:", numero)
	fmt.Println("Endereço de numero:", &numero)
	fmt.Println("Valor do ponteiro:", ponteiro)
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiro)

	*ponteiro = 20
	fmt.Println("Novo valor de numero após alteração via ponteiro:", numero)
}
