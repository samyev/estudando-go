package main

// o go utiliza bibliotecas, e para importar uma biblioteca eu uso o import
import "fmt" // aqui eu importo o pacote fmt

func main() {
	fmt.Println("Hello, Wolrd!")

	// ao rodar o código ele vai mostrar se houve algum erro e o numero de bytes da variavel
	numerobytes, erros := fmt.Println("aqui eu to declarando 2 variaveis", 10, "a baixo eu chamo elas")
	fmt.Println(numerobytes, erros)

	_, erros := fmt.Println("aqui eu to declarando uma variavel anonima")
	fmt.Println(erros)

	a := 10
	b := "Hello"
	c := true

	fmt.Println(a, b, c)
}
