package mediana

import "testing"

//Testes por tabela
func TestMediana(t *testing.T) {
	var listaPar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}
	var listaImpar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7}

	testMediana := map[int]struct {
		listas   Lista
		esperado float64
	}{
		1: {listaPar, 6.5},
		2: {listaImpar, 5},
	}

	for name, tt := range testMediana {
		entrada := tt.listas.Mediana()

		if entrada != tt.esperado {
			t.Errorf("%v: entrada %v, esperado %v", name, entrada, tt.esperado)
		}
	}
}
