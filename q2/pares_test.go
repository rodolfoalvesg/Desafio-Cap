package pares

import (
	"testing"
)

func TestPares(t *testing.T) {
	var listA = []int{1, 5, 3, 4, 2}
	var listB = []int{3, 7, 4, 5, 12, 20, 16, 13, 2, 17, 9, 10}
	var listC = []int{2, 6, 7, 9, 1, 13, 15}

	testPares := map[int]struct {
		list Pares
		want int
	}{
		1: {Pares{X: 2, N: listA}, 3},
		2: {Pares{X: 4, N: listB}, 10},
		3: {Pares{X: 3, N: listC}, 8},
	}

	for indices, tt := range testPares {
		entries := tt.list.EncontraPares()
		if entries != tt.want {
			t.Errorf("%v: entrie %v, want %v", indices, entries, tt.want)
		}
	}

}
