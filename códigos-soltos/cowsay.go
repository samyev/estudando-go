package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("cowsay") // a variável cmd está recebendo o comando cowsay a partir de exec.Command
	// cmd.Stdin receber o input e strings.NewReader ajuda a forncer uma string como entrada para o comando cowsay
	cmd.Stdin = strings.NewReader("Olá galerinha, tudo bom?")
	cmd.Stdout = os.Stdout // cmd.Stdout recebe o output do comando que será exibido no terminal
	cmd.Run()              // cmd.Run executa o comando cowsay
}
