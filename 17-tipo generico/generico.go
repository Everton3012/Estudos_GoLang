package main

import "fmt"

func generico(insert interface{}) {
	fmt.Println("Valor inserido:", insert)
}

func main() {
	fmt.Println("tipos genéricos em Go")

	generico(10)
	generico("Olá, mundo!")
	generico(3.14)

	mapGenerico := map[string]interface{}{
		"nome":      "João",
		"sobrenome": "Silva",
		"idade":     30,
	}
	generico(mapGenerico)
}
