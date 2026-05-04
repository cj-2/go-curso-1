package main

import (
	"fmt"
	"nome_modulo/auxiliar" // nome do modulo "/"  nome do pacote

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste normal mesmo.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("tesssste.com")
	fmt.Println(erro.Error())
}
