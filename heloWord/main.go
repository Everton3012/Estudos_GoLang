package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("hello world")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")

	fmt.Println(erro)
}
