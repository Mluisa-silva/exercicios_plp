package main

import "fmt"

type Carro struct {
	marca     string
	modelo    string
	ano       int
	velocidade int
}

func (c *Carro) Acelerar(incremento int) {
	if incremento > 0 {
		c.velocidade += incremento
		fmt.Printf("O carro acelerou %d km/h.\n", incremento)
	} else {
		fmt.Println("O valor de aceleração deve ser positivo.")
	}
}

func (c *Carro) Frear(decremento int) {
	if decremento > 0 && c.velocidade-decremento >= 0 {
		c.velocidade -= decremento
		fmt.Printf("O carro freou %d km/h.\n", decremento)
	} else {
		fmt.Println("A frenagem não pode resultar em velocidade negativa.")
	}
}

func (c *Carro) ExibirVelocidade() {
	fmt.Printf("Velocidade atual: %d km/h\n", c.velocidade)
}

func (c *Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s\n", c.marca)
	fmt.Printf("Modelo: %s\n", c.modelo)
	fmt.Printf("Ano: %d\n", c.ano)
}
