package main

import "fmt"

// não delcare variaveis se não for usar, go não aceita, ele declara como lixo
var a = "aqui eu to declarando uma variavel fora de um bloco"
var x = "pra declarar uma variavel fora de um bloco eu uso ="

func main() {
	fmt.Println(a)
	fmt.Println(x)

	marmota := "isso é um declarador curto, só funciona dentro de um code block"
	fmt.Println(marmota)

	y := 10
	x := "Bom dia!"
	z := 2.3
	// aqui eu to chamando o z novamente mais declarando uma nova variável, então o valor de z muda
	w, z := 3, 4

	// %v é o valor da variavel, %T é o tipo da variavel
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	fmt.Printf("w: %v, %T\n", w, w)

	// aqui eu uso operadores aritimeticos
	e := 10 + 10
	f := 12 - 2
	g := 5 * 10
	h := 10 / 2

	fmt.Println(e, f, g, h)
}