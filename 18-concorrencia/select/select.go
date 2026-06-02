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
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	// for {
	// 	mensagamCanal1 := <-canal1
	// 	fmt.Println(mensagamCanal1)

	// 	mensagamCanal2 := <-canal2
	// 	fmt.Println(mensagamCanal2)
	// }

	for {
		select {
		case mensagamCanal1 := <-canal1:
			fmt.Println(mensagamCanal1)
		case mensagamCanal2 := <-canal2:
			fmt.Println(mensagamCanal2)
		}
	}

}
