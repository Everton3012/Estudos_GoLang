package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) falar() {
	fmt.Println("Olá, meu nome é", p.nome, "e tenho", p.idade, "anos.")
}

func (p *Pessoa) salvar() {
	fmt.Printf("Salvando os dados de : %s", p.nome)
}

func main() {
	fmt.Println("Metodos em Go")
	p := Pessoa{nome: "Alice", idade: 30}
	p.falar()
	p.salvar()
}
