package main

import "testing"

func TestOla(t *testing.T) {
	esperado := Ola()
	resultado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
