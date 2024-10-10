package main

import (
    "fmt"
)

type ContaBancaria struct {
    titular string
    saldo   float64
}


func NovoContaBancaria(titular string, saldoInicial float64) *ContaBancaria {
    return &ContaBancaria{
        titular: titular,
        saldo:   saldoInicial,
    }
}

func (c *ContaBancaria) Depositar(valor float64) {
    if valor > 0 {
        c.saldo += valor
        fmt.Printf("Depósito de R$ %.2f realizado com sucesso.\n", valor)
    } else {
        fmt.Println("O valor do depósito deve ser positivo.")
    }
}

func (c *ContaBancaria) Sacar(valor float64) {
    if valor > 0 && valor <= c.saldo {
        c.saldo -= valor
        fmt.Printf("Saque de R$ %.2f realizado com sucesso.\n", valor)
    } else if valor > c.saldo {
        fmt.Println("Saldo insuficiente para realizar o saque.")
    } else {
        fmt.Println("O valor do saque deve ser positivo.")
    }
}

func (c *ContaBancaria) GetSaldo() float64 {
    return c.saldo
}

func (c *ContaBancaria) GetTitular() string {
    return c.titular
}
