package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // Array - Tamanho fixo
	s1 := []int{1, 2, 3}  // Slice - Tamanho variável
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um Array! Slice define um pedaço de um array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // Novo slice apontando para o mesmo array
	fmt.Println(a2, s3)

	s4 := s2[:1] // Slice de um slice
	fmt.Println(s2, s4)
}
