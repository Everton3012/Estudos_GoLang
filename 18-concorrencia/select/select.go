package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Mensagem do canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Mensagem do canal 2"
		}
	}()

	for {
		select {
		case mensagem := <-canal1:
			fmt.Println(mensagem)
		case mensagem := <-canal2:
			fmt.Println(mensagem)
		}
	}

}
