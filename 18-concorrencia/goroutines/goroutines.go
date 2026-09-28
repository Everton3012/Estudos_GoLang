package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines em Go")

	go escrever("Escrevendo...")

	escrever("Fim do programa")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
