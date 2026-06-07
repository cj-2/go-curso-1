package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Major", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada das Rosas", "Estrada"},
		{"Rodovia Jussara", "Rodovia"},
		// {"Praça da Chuva", "Tipo Inválido"},
		{"RUA DA LUA", "Rua"},
		// {"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		resultadoTipoEndereco := TipoDeEndereco(cenario.enderecoInserido)

		if resultadoTipoEndereco != cenario.retornoEsperado {
			t.Errorf("Endereço diferente do esperado. Esperava %s e recebeu %s", cenario.retornoEsperado, resultadoTipoEndereco)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if (2 * 3) == 9 {
		t.Errorf("Errou a matemática")
	}
}
