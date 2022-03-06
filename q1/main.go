package main

import (
	"fmt"

	"q1/mediana"
)

func main() {
	listaPar := &mediana.Lista{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}
	listaImpar := &mediana.Lista{9, 2, 1, 4, 6, 8, 3, 5, 7}

	medianaPar := listaPar.Mediana()
	medianaImpar := listaImpar.Mediana()

	fmt.Println(medianaPar, " ", medianaImpar)
}
