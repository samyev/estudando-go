package main

import "fmt"

type hotdog int
var a hotdog = 101

func main() {
	x := 10
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("x: %v, %T\n", x, x)

	// convertendo o valor de x para o tipo hotdog, recebendo o valor de a
	x = int(a)
	fmt.Printf("x: %v, %T\n", x, x)
}