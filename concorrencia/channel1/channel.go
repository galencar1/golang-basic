package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // Enviando dados para um canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
