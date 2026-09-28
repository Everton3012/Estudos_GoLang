package main

import "fmt"

func clousure() func() {
	x := 10

	return func() {
		fmt.Println("Valor de x:", x)
	}
}

func main() {
	fmt.Println("closure em Go")

	texto := "Olá, mundo!"
	fmt.Println(texto)
	f := clousure()
	f()
}
