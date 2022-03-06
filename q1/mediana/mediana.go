package mediana

import (
	"sort"
)

type Lista []int

//Mediana() recebe uma slice(lista) através do tipo 'Lista'

func (l Lista) Mediana() float64 {
	sort.Ints(l)       //Organiza a lista em ordem crescente
	tamLista := len(l) //Salva o tamanho da Lista

	//Se o tamanho da lista for par é obtida a média dos dois valores centrais.
	if tamLista%2 == 0 {
		metade := tamLista / 2
		return float64((l[metade] + l[metade+1])) / float64(2)
	}

	//Caso o tamanho da lista seja impar retorna o valor central
	return float64(l[tamLista/2])
}
