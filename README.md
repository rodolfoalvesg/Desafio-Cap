# Desafio-Cap
Exercício proposta na Lista da CAP.

##  Tecnologias e Recursos necessários para exercução
 - Golang version go1.17.8
 - Terminal ou uma IDE

## Comandos necessários para execução
- Primeiramente, faz-se necessário que você esteja na pacote/diretório a qual se deseja executar.
   - Execução Normal : `go run nomedoarquivo.go` 
   - Execução do Teste: `go test`
   - Execução de dimensão do Teste: `go test -cover`

##  Lista de Questões
### Questão 01

A mediana de uma lista de números é basicamente o elemento que se encontra no meio da lista após a ordenação. Dada uma lista de números com um número ímpar de elementos, desenvolva um algoritmo que encontre a mediana.

Exemplo: 
* Entrada: `arr = [9, 2, 1, 4, 6]`
* Saída: `4`

### Questão 01 - Resolução

#### Referência do Package: `q1/`
 Arquivos: 
 - `mediana.go`:
      - Temos o método principal `Mediada()` que é implementado por uma `Lista` do tipo Slice.
      - Ao ser executada, `Lista` é ordenada pelo método `sort.Ints()`
      - Em seguida obtém-se o tamanho da lista.
      - A lista atendendo ao critério, é realizada operação de cálculo e retorna o valor da mediana.
     
 - `mediana_test.go`
      - Aqui temos a implementação de testes por TDD (test Driven Development)
      - Lista contém a quantidade de elementos PAR: `var listaPar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}`
      - Podemos passar a lista com os valores de testes, junto ao valor esperado, conforme formato: `{listaPar, 6.5}`

Entradas Utilizadas: 
 - `var listaPar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7, 10}`
 - `var listaImpar = []int{9, 2, 1, 4, 6, 8, 3, 5, 7}`

Saídas Obtidas: 
 - Medianas: `6.5` e `5` respectivamente.
