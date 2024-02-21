package main

import (
	"fmt"
	"sync"
)

func minhaGoroutine(wg *sync.WaitGroup) {
	defer wg.Done() // Marca a goroutine como concluída quando terminar
	fmt.Println("Executando minha goroutine...")
}

func main() {
	var wg sync.WaitGroup

	// Adiciona uma goroutine ao WaitGroup
	wg.Add(1)

	// Inicia uma goroutine
	go minhaGoroutine(&wg)

	// Espera todas as goroutines serem concluídas
	wg.Wait()

	fmt.Println("Programa principal encerrado.")
}
