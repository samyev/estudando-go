package main

import (
	"fmt"
)

func main() {
	fmt.Println("Escolha uma opção:\n")

	var op int
	fmt.Scanln(&op)

	var input float64
	fmt.Println("Digite a temperatura:")
	fmt.Scanln(&input)

	var temperatura, unit string

	switch op {
	case 1:
		temperatura, unit = "Celsius", "Fahrenheit"
	case 2:
		temperatura, unit = "Celsius", "Kelvin"
	case 3:
		temperatura, unit = "Fahrenheit", "Celsius"
	case 4:
		temperatura, unit = "Fahrenheit", "Kelvin"
	case 5:
		temperatura, unit = "Kelvin", "Celsius"
	case 6:
		temperatura, unit = "Kelvin", "Fahrenheit"
	default:
		fmt.Println("Opção inválida")
		return
	}

	result := convert(input, temperatura, unit)
	fmt.Printf("%.2f %s = %.2f %s\n", input, temperatura, result, unit)

}

func convert(t float64, temperatura string, unit string) float64 {
	switch temperatura {
	case "Celsius":
		switch unit {
		case "Fahrenheit":
			return (t * 9 / 5) + 32
		case "Kelvin":
			return t + 273.15
		}
	case "Fahrenheit":
		switch unit {
		case "Celsius":
			return (t - 32) * 5 / 9
		case "Kelvin":
			return (t + 459.67) * 5 / 9
		}
	case "Kelvin":
		switch unit {
		case "Celsius":
			return t - 273.15
		case "Fahrenheit":
			return (t * 9 / 5) - 459.67
		}
	}
	return t
}
