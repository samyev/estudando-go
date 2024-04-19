package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) DentroDaFaixaEtaria(min, max int) bool {
	return p.idade >= min && p.idade <= max
}

func FiltrarFaixaEtaria(pessoas []Pessoa, min, max int) []Pessoa {
	var pessoasNaFaixa []Pessoa
	for _, pessoa := range pessoas {
		if pessoa.DentroDaFaixaEtaria(min, max) {
			pessoasNaFaixa = append(pessoasNaFaixa, pessoa)
		}
	}
	return pessoasNaFaixa
}

func main() {
	s := []Pessoa{{"Joe", 18}, {"Hana", 25}, {"Tomas", 30}, {"Thor", 40}}

	filtro := FiltrarFaixaEtaria(s, 18, 30)

	fmt.Printf("Pessoas na faixa etária de 18-30 anos:\n")
	for _, pessoa := range filtro {
		fmt.Println(pessoa.nome)
	}
}
