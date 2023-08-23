package main

import "fmt"

func main() {
	// literal := `"aqui eu to declarando\n uma raw \"string\" literals"`
	// strings := "\naqui eu to declarando\n uma \"string\" literals"

	// fmt.Println(literal)
	// fmt.Println(strings)

	x := "ola"
	y := "hello"

	z := fmt.Sprint(x, y)
	fmt.Println(z)
}