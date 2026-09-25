package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	var numeroReal1 float64 = 1.1

	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12222.1111

	fmt.Println(numeroReal2)

	char := 'B'
	fmt.Println(char)

	var numero3 int16
	fmt.Println(numero3)

	var texto string
	fmt.Println(texto)
	fmt.Println(texto)

	var boolean1 bool = true

	fmt.Println(boolean1)

	var boolean2 bool = false

	fmt.Println(boolean2)

	var erro error

	fmt.Println(erro)

	var erro2 error = errors.New("Erro 2")
	fmt.Println(erro2)

}
