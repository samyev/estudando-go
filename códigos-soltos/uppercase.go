package main

import (
	"fmt"
	"strings"
)

const name string = "meu nome é samylle"

func main() {
	var a int
	var b int

	a = 20
	b = 4
	soma := a + b

	fmt.Printf("\n %s e eu tenho %d anos \n", name, soma)

	search := strings.Contains(name, "s")
	fmt.Printf("\n A frase '%s' contém a letra s?\n Resposta: %t\n\n", name, search)

	uppercase := strings.ToUpper(name)
	fmt.Println("\n", uppercase)

}
