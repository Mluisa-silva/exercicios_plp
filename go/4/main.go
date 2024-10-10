package main

import (
	"fmt"
)

type Motor struct {
	Potencia int
	Tipo     string
}

func (m Motor) ExibirInformacoes() {
	fmt.Printf("Potência do motor: %d CV\n", m.Potencia)
	fmt.Printf("Tipo de motor: %s\n", m.Tipo)
}

type Carro struct {
	Marca  string
	Modelo string
	Motor  Motor 
}

func (c Carro) ExibirInformacoes() {
	fmt.Printf("Marca: %s\n", c.Marca)
	fmt.Printf("Modelo: %s\n", c.Modelo)
	c.Motor.ExibirInformacoes()
}

func main() {
	
	motor1 := Motor{Potencia: 150, Tipo: "Gasolina"}

	carro1 := Carro{Marca: "Toyota", Modelo: "Corolla", Motor: motor1}

	carro1.ExibirInformacoes()
}
