package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "variavel 1"
	fmt.Println(variavel1)
	variavel2 := "variavel 2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "constante1"

	fmt.Println(constante1)

	const (
		constante2 string = "constante2"
	)

	fmt.Println(constante2)
}
