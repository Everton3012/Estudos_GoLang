package main

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
	altura    float64
}

type Estudante struct {
	Pessoa
	curso string
}

func main() {
	estudante := Estudante{
		Pessoa: Pessoa{
			nome:      "João",
			sobrenome: "Silva",
			idade:     20,
			altura:    1.75,
		},
		curso: "Engenharia",
	}

	println(estudante.nome)
	println(estudante.sobrenome)
	println(estudante.idade)
	println(estudante.altura)
	println(estudante.curso)
}
