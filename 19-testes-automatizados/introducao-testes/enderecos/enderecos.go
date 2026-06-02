package enderecos

import (
	"slices"
	"strings"
)

// TipoDeEndereco Retorna se é um tipo válido de endereço.
func TipoDeEndereco(endereco string) string {
	endereco = strings.ToLower(endereco)
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavaraEndereco := strings.Split(endereco, " ")[0]

	enderecoTemTipoValido := false

	// for _, tipo := range tiposValidos {
	// 	if tipo == primeiraPalavaraEndereco {
	// 		enderecoTemTipoValido = true
	// 		break
	// 	}
	// }

	enderecoTemTipoValido = slices.Contains(tiposValidos, primeiraPalavaraEndereco)

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavaraEndereco)
	}

	return "Tipo Inválido"
}
