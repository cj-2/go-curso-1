package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Hellow!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go.")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
