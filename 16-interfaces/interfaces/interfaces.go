package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (c retangulo) area() float64 {
	return c.altura * c.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	// return (c.raio * c.raio) * math.Pi
	return math.Pow(c.raio, 2) * math.Pi
}

func escreverArea(f forma) {
	fmt.Println("Area =>", f.area())
}

func main() {
	r := retangulo{12, 12}
	escreverArea(r)

	c := circulo{12}
	escreverArea(c)
}
