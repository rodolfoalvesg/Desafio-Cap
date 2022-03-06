# Desafio-Cap
Exercício proposta na Lista da CAP.

##  Tecnologia utilizada
 - Golang


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
      - Aqui temos a implementação do teste por TDD, de modo a otimizar nossos testes com maior número de variáveis de possibilidades.
      - 

