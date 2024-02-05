// fetch exibe o conteúdo encontrado em cada URL especificada.

package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") { // ex. 1.8 Add strings.HasPrefix()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err, resp.Status) // ex. 1.9 Add resp.Status
				os.Exit(1)
			}
			// b, err := ioutil.ReadAll(resp.Body)
			b, err := io.Copy(os.Stdout ,resp.Body) // ex. 1.7 - Add io.Copy
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b, resp.Status) // ex. 1.9 Add resp.Status
		} else {
			fmt.Fprintf(os.Stderr, "fetch: URL '%s' não possui o prefixo 'https://' ou 'http://'\n", url)
			os.Exit(1)
		}
	}
}
