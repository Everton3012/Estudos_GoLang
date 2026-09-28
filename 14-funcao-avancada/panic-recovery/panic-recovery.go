package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada do panic:", r)
	}
}

func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6, o que não é permitido.")
}

func main() {
	fmt.Println("panic-recovery em Go")
	aluno1 := alunoAprovado(7, 8)
	fmt.Println("Aluno 1 aprovado:", aluno1)

	aluno2 := alunoAprovado(5, 4)
	fmt.Println("Aluno 2 aprovado:", aluno2)

	aluno3 := alunoAprovado(6, 6)
	fmt.Println("Aluno 3 aprovado:", aluno3)
}
