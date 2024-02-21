package main

import (
	"fmt"
	"time"
)

func minhaGoroutine(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Mensagem %d", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

func main() {
	// Cria um canal
	c := make(chan string)

	// Inicia uma goroutine que envia mensagens para o canal
	go minhaGoroutine(c)

	// Lê as mensagens do canal e imprime
	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Println("Programa principal encerrado.")
}
