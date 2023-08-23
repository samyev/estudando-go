package main

import "fmt"

var ano = 2023

func main() {
	fmt.Println(ano)

	qualquercoisa()
}

func qualquercoisa() {
	mes := 8

	fmt.Println(ano, mes)
	fmt.Printf("ano %v/mes %v\n", ano, mes)
	
}