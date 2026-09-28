package main

import "fmt"

var n int

func init() {
	fmt.Println("Inicialização em Go")
	n = 10
}

func main() {

	fmt.Println("inicializado e pronto para executar")
	fmt.Println("Valor de n:", n)
	fmt.Println("Valor de n após incremento:", n+5)
}
