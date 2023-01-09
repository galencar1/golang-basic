package main

import "fmt"

func main() {
	//inicializando map de modo literal
	funcsESalarios := map[string]float64{
		"Jose Joao":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}
	//adicionando novos valores no map
	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // Caso tente deletar um valor inexistente, não acontecerá nada.

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
