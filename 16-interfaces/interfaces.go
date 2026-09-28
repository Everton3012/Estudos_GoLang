package main

import "fmt"

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return 3.14 * c.raio * c.raio
}

func escreverArea(f forma) {
	fmt.Println("Área:", f.area())
}

func main() {
	fmt.Println("Interfaces em Go")
	r := retangulo{altura: 10, largura: 5}
	c := circulo{raio: 7}

	escreverArea(r)
	escreverArea(c)
}
