package main

import (
	"fmt"
)

type Usuario struct {
	nome     string
	idade    int
	endereco Endereco
}

type Endereco struct {
	rua    string
	numero int
}

func main() {
	fmt.Println("Structs in Go")
	Usuario1 := Usuario{"João", 30, Endereco{"Rua A", 123}}
	Usuario2 := Usuario{idade: 30}
	fmt.Println(Usuario2)
	fmt.Println(Usuario1)
	fmt.Println(Usuario1.nome)
	fmt.Println(Usuario1.idade)
	fmt.Println(Usuario1.endereco.rua)
}
