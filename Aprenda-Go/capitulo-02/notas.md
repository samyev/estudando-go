# notas sobre golang :=

- go não aceita variáveis sem uso
- se utiliza _ quando se quer declarar uma variável anonima
- := isso é um declarador curto, só pode ser usado em blocos de código, faz declaração e atribuição
- só se chama uma variável já declarada se for declarar uma nova variável (ex: nova, antiga := 1, 2)
- isso = faz uma declaração de um valor pra uma variável que está fora de um bloco, ela tem um escopo de uso abrangente em qualquer local dentro de um package, diferente do declarador curto, que só é válido apartir de sua declaração
- keywords: são palavras reservadas
- statement (declaração, afirmação) → é uma linha de código formada por expressões, expressa uma ação a ser executada.  
- expressão → qualquer coisa que "produz um resultado" 
- ao encontrar "..." quer dizer que isto é uma função variática, ou seja você pode declarar valores de qualquer tipo.
- tipos são imutáveis, são fixos.
- em go só existe o tipo converstion
- valor zero, é o valor que o computador atribui quando você não define valor na variável