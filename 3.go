package main

import "fmt"

type Produto struct {
	nome       string
	preco      float64
	quantidade int
}

func mudar(p *Produto, valor float64) {
	p.preco = valor
}

func main() {
	var valor float64
	fmt.Println("Escreva o novo valor: ")
	fmt.Scanln(&valor)
	/*var valorpn *float64 = &valor*/
	p := Produto{nome: "Sanderson", preco: 526, quantidade: 5}
	fmt.Println("Struct original: ", p)
	mudar(&p, valor)
	fmt.Println("Novo preço: ", p.preco)
	fmt.Println("Structs atualizada", p)
}
