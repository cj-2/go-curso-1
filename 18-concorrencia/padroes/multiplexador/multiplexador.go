package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case m1 := <-canal1:
				canalSaida <- m1
			case m2 := <-canal2:
				canalSaida <- m2
			}
		}
	}()

	return canalSaida
}

func main() {
	canal := multiplexar(escrever("Batatinha"), escrever("Hamburger"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
