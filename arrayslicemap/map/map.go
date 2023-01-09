package main

import "fmt"

func main() {
	//var aprovados map[int]string // map com chave int e valor string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678900] = "Maria"
	aprovados[12345678901] = "Pedro"
	aprovados[12345678902] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	//para consultar um valor do map
	fmt.Println(aprovados[12345678902])

	//para deletar um valor do map
	delete(aprovados, 12345678902)
	fmt.Println(aprovados)
}
