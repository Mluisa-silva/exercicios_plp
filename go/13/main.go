package main

import (
	"errors"
	"fmt"
)

type SaldoInsuficienteException struct {
	mensagem string
}

func (e *SaldoInsuficienteException) Error() string {
	return e.mensagem
}

type ContaBancaria struct {
	saldo float64
}

func NewContaBancaria(saldoInicial float64) *ContaBancaria {
	return &ContaBancaria{saldo: saldoInicial}
}

func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.saldo {
		return &SaldoInsuficienteException{"Saldo insuficiente para realizar o saque."}
	}
	c.saldo -= valor
	return nil
}

func (c *ContaBancaria) GetSaldo() float64 {
	return c.saldo
}

func main() {
	conta := NewContaBancaria(100.0)
	err := conta.Sacar(150.0)
	if err != nil {
		fmt.Println(err)
	}
}
