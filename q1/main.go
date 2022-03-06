package main

import (
	"fmt"

	"github.com/rodolfoalvesg/Desafio-Cap/q1/mediana"
)

func main() {
	listaPar := &mediana.Lista{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}

	mediana := listaPar.Mediana()

	fmt.Println(mediana)
}
