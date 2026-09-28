package main

import "fmt"

func inverterSinal(numero *int) int {
	*numero = *numero * -1
	return *numero
}

func main() {
	fmt.Println("funcão com ponteiros em Go")
	numero := 10
	fmt.Println("Número original:", numero)
	inverterSinal(&numero)
	fmt.Println("Número após inverter o sinal:", numero)

}
