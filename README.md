# Desafio-Cap
Exercício proposta na Lista da CAP.

##  Tecnologias e Recursos necessários para execução
 - Golang version go1.17.8
 - Terminal ou uma IDE

## Comandos necessários para execução

- Entre no diretório da questão desejada para Execução Normal : `go run main.go`
- Para testes, faz-se necessário que você esteja na pacote/diretório a qual se deseja executar. 
   - Execução do Teste: `go test`
   - Execução de dimensão do Teste: `go test -cover`

##  Lista de Questões
### Questão 01

A mediana de uma lista de números é basicamente o elemento que se encontra no meio da lista após a ordenação. Dada uma lista de números com um número ímpar de elementos, desenvolva um algoritmo que encontre a mediana.

Exemplo: 
* Entrada: `arr = [9, 2, 1, 4, 6]`
* Saída: `4`

### Questão 01 - Resolução

#### Referência do Package: `q1/mediana`
Arquivos: 
 - `mediana.go`:
 - `mediana_test.go`

Métodos Principais:
- `Mediana()`
 
Entradas Utilizadas: 
 - `var listaPar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}`
 - `var listaImpar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7}`

Saídas Obtidas: 
 - Medianas: `6.5` e `5` respectivamente.


### Questão 02
Dado um vetor de inteiros n e um inteiro qualquer x. Construa um algoritmo que determine o número de elementos pares do vetor que tem uma diferença igual ao valor de x.

Exemplo:
- Entrada: `n = [1, 5, 3, 4, 2]`
- Saída: `3`

Explicação:
- Existem 3 pares de inteiros no vetor com uma diferença de 2: `[5, 3], [4, 2] e [3, 1]`.

### Questão 02 - Resolução
#### Referência do Package: `q2/pares`

Arquivos: 
 - `pares.go`
 - `pares_test.go`

Métodos Principais:
- `EncontraPares()`
- `validaPares()`
 
Entradas Utilizadas: 
 - `var listA = []int{1, 5, 3, 4, 2}`
	- `var listB = []int{3, 7, 4, 5, 12, 20, 16, 13, 2, 17, 9, 10}`
	- `var listC = []int{2, 6, 7, 9, 1, 13, 15}`

Saídas Obtidas: 
 - Pares: `3`, `10` e `8` respectivamente.

### Questão 03
Um texto precisa ser encriptado usando o seguinte esquema. Primeiro, os espaços são removidos do texto. Então, os caracteres são escritos em um grid, no qual as linhas e colunas tem as seguintes regras:

`sqrt(T)<=linha<=coluna<=sqrt(T)`

- Considere T, como o tamanho do texto.
- Se certifique de que linhas x colunas >= .
- Se múltiplos grids satisfazem as condições, escolha aquele com a menor área.
- Escreva um algoritmo que ao receber uma string s, mostre a mensagem encriptada de acordo com as regras descritas.

Exemplo:

- Entrada: `s = tenha um bom dia`
- Saída: `taoa eum nmd hbi`

### Questão 03 - Resolução
#### Referência do Package: `q3/cripto`

Arquivos: 
 - `cripto.go`
 - `cripto_test.go`

Métodos Principais:
- `Cripto()`
- `criarGrid()`
- `criarMensagemEncriptada()`
 
Entradas Utilizadas: 
 - `var fraseA cripto.String = "tenha um bom dia"`
	- `var fraseB cripto.String = "ola mundo"`

Saídas Obtidas: 
 - Mensagens: `taoa eum nmd hbi` e `omd luo an` respectivamente.
