package main

import "fmt"

type imprivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.99}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	imprimir(produto{"Camisa Branca", 29.99})
	imprimir(pessoa{"Noah", "Miguel"})
}
