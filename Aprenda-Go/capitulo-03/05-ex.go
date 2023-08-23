package main

import "fmt"

type batatinha int
var x batatinha
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Printf("%v\n", x)

	y = int(x)

	fmt.Printf("%T\n", y)
	fmt.Println(y)
}