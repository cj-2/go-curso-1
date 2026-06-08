package formas_test

import (
	"math"
	. "teste-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area incorreta. Esperada: %f; Recebida: %f;", areaEsperada, areaRecebida)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		cir := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area incorreta. Esperada: %f; Recebida: %f;", areaEsperada, areaRecebida)
		}
	})
}
