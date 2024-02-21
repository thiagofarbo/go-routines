package main

import (
	"fmt"
	"time"
)

func minhaGoroutine(numero int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", numero, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Inicia duas goroutines
	go minhaGoroutine(1)
	go minhaGoroutine(2)

	// Espera um pouco para que as goroutines tenham tempo de executar
	time.Sleep(1 * time.Second)

	fmt.Println("Programa principal encerrado.")
}
