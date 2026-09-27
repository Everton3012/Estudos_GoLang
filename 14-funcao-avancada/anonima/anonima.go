package main

import "fmt"

func main() {
	funcaoAnonima := func() {
		fmt.Println("Função Anônima")
	}

	funcaoAnonima()

	func(texto string) {
		fmt.Println("Função Anônima Executada", texto)
	}("com parâmetro")

	func() {
		fmt.Println("Função Anônima Executada sem parâmetro")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Função Anônima Executada com retorno: %s", texto)
	}("com parâmetro")
	fmt.Println(retorno)
}
