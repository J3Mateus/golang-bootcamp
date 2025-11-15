package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
	perimetro() float64
}

type Retangulo struct {
	base   float64
	altura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) area() float64 {
	return r.base * r.altura
}

func (r Retangulo) perimetro() float64 {
	return 2 * (r.base + r.altura)
}

func (r Circulo) perimetro() float64 {
	return 2 * math.Pi * r.raio
}

func (r Circulo) area() float64 {
	return math.Pi * r.raio * r.raio
}

func main() {
	retangulo := Retangulo{base: 5, altura: 10}
	circulo := Circulo{raio: 7}

	formas := []Forma{retangulo, circulo}

	for _, forma := range formas {
		fmt.Printf("--- %T ---\n", forma)
		fmt.Printf("Área: %.2f\n", forma.area())
		fmt.Printf("Perímetro: %.2f\n", forma.perimetro())
	}
}
