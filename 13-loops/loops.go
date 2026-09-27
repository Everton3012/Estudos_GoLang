package main

import "fmt"

func main() {
	fmt.Println("loops em Go")

	// Loop for
	for i := 0; i < 5; i++ {
		fmt.Println("Loop for:", i)
	}

	// Loop while
	j := 0
	for j < 5 {
		fmt.Println("Loop while:", j)
		j++
	}

	// Loop do-while (simulado)
	k := 0
	for {
		fmt.Println("Loop do-while:", k)
		k++
		if k >= 5 {
			break
		}
	}

	nomes := []string{"João", "Maria", "Pedro"}

	for _, nome := range nomes {
		fmt.Println("Loop range:", nome)
	}

	for i, letra := range "GoLang" {
		fmt.Println("Loop range letra:", i+1, string(letra))
	}

	usuarios := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     "30",
	}

	for chave, valor := range usuarios {
		fmt.Println("Loop range map:", chave, valor)
	}
}
