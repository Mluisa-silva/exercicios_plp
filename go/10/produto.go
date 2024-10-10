package main

import (
    "fmt"
)

type Produto struct {
    Nome  string
    Preco float64
}

func (p Produto) Soma(outro Produto) Produto {
    return Produto{
        Nome:  p.Nome + " + " + outro.Nome,
        Preco: p.Preco + outro.Preco,
    }
}

func (p Produto) String() string {
    return fmt.Sprintf("%s: R$%.2f", p.Nome, p.Preco)
}

