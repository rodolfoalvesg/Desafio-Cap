package main

import (
	"fmt"
	"q2/pares"
)

func main() {
	listaA := &pares.Pares{X: 2, N: []int{1, 5, 3, 4, 2}}
	listaB := &pares.Pares{X: 4, N: []int{3, 7, 4, 5, 12, 20, 16, 13, 2, 17, 9, 10}}
	listaC := &pares.Pares{X: 3, N: []int{2, 6, 7, 9, 1, 13, 15}}

	listaParA := listaA.EncontraPares()
	listaParB := listaB.EncontraPares()
	listaParC := listaC.EncontraPares()

	fmt.Printf("%v, %v e %v\n", listaParA, listaParB, listaParC)

}
