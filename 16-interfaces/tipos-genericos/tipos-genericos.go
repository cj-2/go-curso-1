package main

import "fmt"

func generica(i interface{}) {
	fmt.Println("Função genérica, valor =>", i)
}

func generica2(i any) {
	fmt.Println("Função genérica, valor =>", i)
}

func main() {
	generica(2)
	generica("Uau!")
	generica(false)

	generica2("Paçoca.")
}
