package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle em Go")

	numero := 10
	if numero > 0 {
		fmt.Println("O número é positivo")
	} else if numero < 0 {
		fmt.Println("O número é negativo")
	} else {
		fmt.Println("O número é zero")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Outro número é positivo")
	} else if outroNumero < 0 {
		fmt.Println("Outro número é negativo")
	} else {
		fmt.Println("Outro número é zero")
	}

	switch numero {
	case 1:
		fmt.Println("O número é um")
	case 2:
		fmt.Println("O número é dois")
	default:
		fmt.Println("O número não é nem um nem dois")
	}
}
