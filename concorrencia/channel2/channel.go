package main

import (
	"fmt"
	"time"
)

// Channel - é a forma de comunicação entre goroutines

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)

}
