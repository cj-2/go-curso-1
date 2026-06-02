package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Major Gote 155")
	fmt.Println(tipoEndereco)
}
