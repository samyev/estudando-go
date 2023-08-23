package main

import "fmt"

type hotdog int
var a hotdog

func main() {
	fmt.Printf("%v, %T\n", a, a)
}