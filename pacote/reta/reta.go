package main

import "math"

//Iniciando  com letra maiúscula é PÚBLICO(visível dentro de fora do PACOTE)
//Iniciando  com letra minuscula é PRIVADO(visível apenas fora do PACOTE)

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Ponto representa uma coordenada no ponto cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distância é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
