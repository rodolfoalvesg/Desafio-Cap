package cripto

import (
	"math"
	"strings"
)

type String string

// criarMensagemEncriptada, recebe a grid e a raiz e distribui os valores conforme posição
func criarMensagemEncriptada(tabela [][]string, raiz int) string {
	mensagem := []string{}              //Slice da mensagem encriptada
	ultimoIndice := len(tabela[raiz-1]) //Obtém o tamanho da última linha

	for pos := 0; pos <= raiz; pos++ {

		//Se o última linha da Grid tiver tamanho menor que as demais, há uma adequação
		if ultimoIndice == 0 {
			raiz--
		}

		//Mapeia as linhas da Grid e adiciona a mensagem conforme condição
		for i := 0; i < raiz; i++ {
			mensagem = append(mensagem, tabela[i][pos])
			if (i + 1) == raiz {
				mensagem = append(mensagem, " ")
			}
		}
		ultimoIndice--

	}

	return strings.Join(mensagem, "")
}

//CriarGrid, cria a estrutura de linhas e colunas simulando uma tabela
func criarGrid(raiz int, texto string) string {
	text := strings.Split(texto, "") //Separa os caracteres no array
	subGrid := []string{}            // Slice unidimensional
	grid := [][]string{}             // Slice multidimensional

	//Responsável por criar a linhas da grid
	for _, char := range text {
		subGrid = append(subGrid, char)
		if len(subGrid) >= raiz {
			grid = append(grid, subGrid)
			subGrid = []string{}
		}
	}

	grid = append(grid, subGrid) // Adiciona a última linha a Grid geral

	return criarMensagemEncriptada(grid, raiz)
}

func (s String) Cripto() string {
	stringUnificada := strings.ReplaceAll(string(s), " ", "") //Retira os espaços da string
	tamString := float64(len(stringUnificada))                //Tamanho da String
	raizString := int(math.RoundToEven(math.Sqrt(tamString))) //Extrai a raiz do Tamanho da String e arredonda

	return criarGrid(raizString, stringUnificada)
}
