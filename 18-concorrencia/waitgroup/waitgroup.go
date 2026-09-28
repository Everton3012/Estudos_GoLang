package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	fmt.Println("Goroutines em Go")

	waitGroup.Add(4)

	go func() {
		escrever("Escrevendo...")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programa em Go 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programa em Go 2")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programa em Go 3")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
