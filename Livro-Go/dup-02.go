// // dup-02 exibe a contagem e o texto das linhas que aparecem mais de 
// // uma vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%v\n%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// NOTA: ignorando erros em potencial de input.Err()
// }

// dup-02 exibe a contagem e o texto das linhas que aparecem mais de
// uma vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	// Mapa para rastrear os arquivos em que cada linha duplicada ocorre
	lineToFileMap := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFileMap, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFileMap, arg)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			files := lineToFileMap[line]
			fmt.Printf("  Arquivos:")
			for _, file := range files {
				fmt.Printf(" %s", file)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFileMap map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		lineToFileMap[line] = append(lineToFileMap[line], filename)
	}
	// NOTA: ignorando erros em potencial de input.Err()
}
