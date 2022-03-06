package main

import (
	"fmt"
	"q3/cripto"
)

func main() {
	var fraseA cripto.String = "tenha um bom dia"
	var fraseB cripto.String = "ola mundo"

	fraseEncriptadaA := fraseA.Cripto()
	fraseEncriptadaB := fraseB.Cripto()

	fmt.Printf("%v | %v\n", fraseEncriptadaA, fraseEncriptadaB)
}
