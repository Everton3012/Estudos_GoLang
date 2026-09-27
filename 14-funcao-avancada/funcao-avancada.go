package main

import "fmt"

func calculosMatematicos(a int, b int) (soma int, subtracao int) {
	soma = a + b
	subtracao = a - b

	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 5)
	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
}
