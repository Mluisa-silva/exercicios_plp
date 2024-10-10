package main

import (
	"fmt"
)

type Funcionario interface {
	calcularSalario() float64
}

type FuncionarioHorista struct {
	nome           string
	cpf            string
	valorHora      float64
	horasTrabalhadas int
}


type FuncionarioAssalariado struct {
	nome        string
	cpf         string
	salarioMensal float64
}


func (f *FuncionarioHorista) calcularSalario() float64 {
	return f.valorHora * float64(f.horasTrabalhadas)
}

func (f *FuncionarioAssalariado) calcularSalario() float64 {
	return f.salarioMensal
}
