package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco - (p.preco * (p.desconto / 100))
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 5,
	}
	fmt.Println(produto1)
	fmt.Printf("Preço com desconto %.2f\n", produto1.precoComDesconto())

	produto2 := produto{"Notebook", 2000.0, 10}
	fmt.Println(produto2)
	fmt.Printf("Preço com desconto: %.2f", produto2.precoComDesconto())
}
