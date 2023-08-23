// 1.2 Modifique o programa echo para exibir o indice e o valor de cada um de seus argumentos, um por linha
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string

	start := time.Now()

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		
		fmt.Println(i, s)
		fmt.Println(start)
	}
}