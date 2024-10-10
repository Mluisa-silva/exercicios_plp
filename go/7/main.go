package main

import "fmt"

type Imprimível interface {
	Imprimir()
}

type Relatório struct {
	conteudo string
}

func (r Relatório) Imprimir() {
	fmt.Println("Imprimindo Relatório:", r.conteudo)
}

type Contrato struct {
	detalhes string
}

func (c Contrato) Imprimir() {
	fmt.Println("Imprimindo Contrato:", c.detalhes)
}

func main() {
	var r Imprimível = Relatório{conteudo: "Relatório Financeiro"}
	var c Imprimível = Contrato{detalhes: "Contrato de Prestação de Serviços"}

	r.Imprimir()
	c.Imprimir()
}
