package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (c Retangulo) Area() float64 {
	return c.Altura * c.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pow(c.Raio, 2) * math.Pi
}
