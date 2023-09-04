package main

import (
	"fmt"
	"strings"
)

type Contato struct {
	Nome, Email string
}

func (d *Contato) alteraEmail(email string){
	d.Email = email
}

func adicionaContato(c Contato, n *[5]Contato){
	for indice, contato := range n {
		if (contato == Contato{}){
			n[indice] = c
			break
		}
	}
}

func main() {
	var pessoas [5]Contato
	var p1 = Contato{
		Nome: "aaa",
		Email: "bbb",
	}

	fmt.Println(pessoas)
	adicionaContato(p1, &pessoas)
	fmt.Println(pessoas)

	x := 5 // declaracao de variavel
	y:= x

	fmt.Println(x, y) // 5 5

	x = 6 // atribuicao de valor
	fmt.Println(x, y) // 

	z := &x // z recebe a referencia de x
	fmt.Println(z) // endereco na memoria em que esta armazenado
	fmt.Println(*z) // deferencia -> retoirna o valor que esta sendo apontado por z

	x = 7
	fmt.Println(x, *z) // 7 7

	var w *int
	fmt.Println(w) // nil -> Nao foi inicializado
	// fmt.Println(*w) -> da um erro!

	mensagem := "Ola, mundo!"
	alteraMensagem(&mensagem)
	fmt.Println(mensagem) // Ola, turma!

	a := &z 

	fmt.Println(z) // endereco de x
	fmt.Println(a) // endereco de z
	fmt.Println(*a) // endereco de x
	fmt.Println(**a) // valor de x
}

func alteraMensagem(msg *string) {
	//strings.ReplaceAll(texto, string, valorProxurado string, valorSubstituir string)
	// Montar uma funcao que substitui "mundo" por "turma"
	*msg = strings.ReplaceAll(*msg, "mundo", "turma")
}

// & referncia o valor da memoria
// * referencia o valor