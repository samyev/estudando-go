# Código Olá Mundo

Aqui você encontra um código básico iniciante de go, que cria uma função que retorna uma segunda função que printa no terminal a frase "Olá, mundo". Em sequida encontramos um segundo código que vai testar se o ola.go realmente retorna o que nós esperamos que ele faça, se sim, ele retorna que código passou no teste, como representado abaixo:

~~~ zsh
PASS
ok      ola     0.002s
~~~

Se não, ele retorna que o código falhou, e printa no terminal o imput e o que ele esperava como output 

~~~ zsh
--- FAIL: TestOla (0.00s)
    ola_test.go:10: resultado 'Olá, mundo!', esperado 'Olá, mundo'
FAIL
exit status 1
FAIL    ola     0.002s
~~~

## Explicação do Código

~~~ go
// ola.go

package main

import "fmt"

func Ola() string {
	return "Olá, mundo"
}

func main() {
	fmt.Println(Ola())
}
~~~

Na linguagem Go há um pacote **main** que define uma função **(func)** main dentro dele.

A palavra reservada **func** é utilizada para que você defina uma função com um nome e um contéudo.

Ao utilizar **import "fmt"**, estamos importando um pacote que contém a função **Println** que será utilizada para imprimir (escrever) um valor no terminal.


~~~ go
// ola_test.go

package main

import "testing"

func TestOla(t *testing.T) {
	esperado := Ola()
	resultado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

~~~

`if`

&nbsp;
Instruções if em Go são muito parecidas com as de outras linguagens.
&nbsp;

`Declarando variáveis`

&nbsp;
Estamos declarando algumas variáveis com a sintaxe nomeDaVariavel := valor, que nos permite reutilizar alguns valores nos nossos testes de maneira legível.
&nbsp;

`t.Errorf`

&nbsp;
O método Errorf é chamado em nosso **t** que irá imprimir uma mensagem e falhar o teste. O sufixo **f** no final de **Errorf** representa que podemos formatar e montar uma string com valores inseridos dentro de valores de preenchimentos **%s**. Quando fazemos um teste falhar, devemos ser bastante claros com o que aconteceu.
