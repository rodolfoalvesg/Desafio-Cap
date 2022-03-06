package pares

import (
	"sort"
)

type Pares struct {
	X int
	N []int
}

//validaPares, recebe o valor da soma de cada sequência e o valor atual e verifica
//se a condição é atendida, retornando um booleano.
func (p *Pares) validaPares(valor int, atual *int) bool {

	if valor%2 == 0 && *atual < valor {
		*atual = valor
		return true
	}
	return false
}

//EncotraPares, possui as variáveis de inicialização, e percorre a lista
//para calculo dos valores.
func (p *Pares) EncontraPares() int {
	somador := 0
	valorAtual := p.X
	acumulador := 0
	sort.Ints(p.N)

	for _, i := range p.N {
		for _, j := range p.N {
			if i != j {
				somador = i + j
				if p.validaPares(somador, &valorAtual) {
					acumulador++
				}
			}

		}
	}

	return acumulador
}
