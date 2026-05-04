package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de inteiros: int8, int16, int32, int64
	var numero int = 100 // int sozinho utiliza a arquitetura do computador como base, ex: int32 ou int64
	numero2 := 100000000000
	fmt.Println(numero, numero2)

	// uint (int sem negativos), tem o mesmo esquema 8, 16, 32, 64

	// alias
	var numero3 rune = 123456 // int32
	var numero4 byte = 123    // uint8
	fmt.Println(numero3, numero4)

	// reais: float32, float64
	var numeroReal1 float32 = 1234.56
	var numeroReal2 float32 = 1234000000000000000.56
	numeroReal3 := 123.456 // aqui vai pela arquitetura do SO, igual apenas "int"
	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	var str1 string = "exemplo 1"
	str2 := "exemplo 2"
	fmt.Println(str1, str2)

	// Go não tem o tipo CHAR
	char := 'B'       // mas caso feito dessas forma, será convertido em um número "rune" referente da tabela ASCI
	fmt.Println(char) // 66

	var texto string // string vazia, valor 0 do tipo
	fmt.Println(texto)

	var booleano1 bool = true
	booleano2 := false
	fmt.Println(booleano1, booleano2)

	// var erro error
	// fmt.Println(erro) // <nil>

	var erro error = errors.New("Erro exemplo.")
	fmt.Println(erro)
}
