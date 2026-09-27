package main

import "fmt"

func main() {
	fmt.Println("Mapas em Go")

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	fmt.Println("Usuário:", usuario["nome"], usuario["sobrenome"], usuario["idade"])

	usuario2 := map[string]map[string]string{
		"joao": {
			"nome":      "João",
			"sobrenome": "Silva",
			"idade":     "30",
		},
		"maria": {
			"nome":      "Maria",
			"sobrenome": "Souza",
			"idade":     "25",
		},
	}
	fmt.Println("Usuário 2:", usuario2["joao"]["nome"], usuario2["joao"]["sobrenome"], usuario2["joao"]["idade"])
	delete(usuario2, "joao")
	fmt.Println("Usuário 2 após delete:", usuario2)

	usuario2["signo"] = map[string]string{
		"nome":      "Signo",
		"sobrenome": "Astrologia",
		"idade":     "N/A",
	}
	fmt.Println("Usuário 2 após adicionar signo:", usuario2)
}
