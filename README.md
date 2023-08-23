# Estudando Golang

## Sobre Golang

A linguagem de programação Go é um projeto de código aberto para tornar os programadores mais produtivos.

Go é expressivo, conciso, limpo e eficiente. Seus mecanismos de simultaneidade facilitam a escrita de programas que tiram o máximo proveito de máquinas multicore e em rede, enquanto seu novo sistema de tipos permite a construção de programas flexíveis e modulares. Go compila rapidamente para código de máquina, mas tem a conveniência da coleta de lixo e o poder da reflexão em tempo de execução. É uma linguagem compilada, de tipagem estática e rápida que parece uma linguagem interpretada e tipada dinamicamente.

## Requisitos de teste

- Versão go: 1.21.0

## Comandos Go
~~~ go
go mod init [packag]
~~~

O comando ```go mod init``` é usado para inicializar um módulo Go no diretório atual ou no diretório especificado pelo argumento [package].

~~~ go
go test
~~~
O comando ```go test``` é usado para executar testes automatizados em seu código Go. Ele procura por funções de teste nos arquivos com o sufixo _test.go dentro do diretório atual e de seus subdiretórios.

~~~ go
go build
~~~
O comando ```go build``` é usado para compilar um ou mais arquivos Go em um único executável ou em um pacote compartilhável (biblioteca).

~~~ go
go mod tidy
~~~
O comando ```go mod tidy``` é usado para ajustar e atualizar o arquivo go.mod para refletir as dependências reais usadas no seu código.

## Referências
* [A Linguagem de Programação Go](https://www.amazon.com.br/Linguagem-Programa%C3%A7%C3%A3o-Go-Alan-Donovan/dp/8575225464/ref=asc_df_8575225464/?tag=googleshopp00-20&linkCode=df0&hvadid=379792215563&hvpos=&hvnetw=g&hvrand=16799462313471420693&hvpone=&hvptwo=&hvqmt=&hvdev=c&hvdvcmdl=&hvlocint=&hvlocphy=1001538&hvtargid=pla-396486666170&psc=1)
* [Aprenda Go](https://www.youtube.com/@AprendaGo/playlists)
* [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)   
* [Docs Golang](https://go.dev/doc/)
