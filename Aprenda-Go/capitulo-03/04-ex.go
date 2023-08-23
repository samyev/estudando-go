package main

import "fmt"

type batatinha int
var x batatinha

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v, %T\n", x, x)
}