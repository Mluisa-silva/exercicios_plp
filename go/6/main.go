package main

import (
    "fmt"
)

type Empregado struct {
    Nome   string
    Cargo  string
    Salario float64
}

func (e Empregado) ExibirDetalhes() {
    fmt.Printf("Nome: %s\n", e.Nome)
    fmt.Printf("Cargo: %s\n", e.Cargo)
    fmt.Printf("Salário: R$ %.2f\n", e.Salario)
}

type Empresa struct {
    Nome      string
    Empregados []Empregado
}

func (emp *Empresa) AdicionarEmpregado(e Empregado) {
    emp.Empregados = append(emp.Empregados, e)
}

func (emp Empresa) ExibirEmpregados() {
    fmt.Printf("Empresa: %s\n", emp.Nome)
    fmt.Println("Lista de Empregados:")
    for _, empregado := range emp.Empregados {
        empregado.ExibirDetalhes()
        fmt.Println("---------------")
    }
}

func main() {
    empresa := Empresa{Nome: "Tech Solutions"}

    emp1 := Empregado{Nome: "João Silva", Cargo: "Desenvolvedor", Salario: 5000.00}
    emp2 := Empregado{Nome: "Maria Oliveira", Cargo: "Gerente de Projetos", Salario: 8000.00}
    emp3 := Empregado{Nome: "Pedro Santos", Cargo: "Analista de Sistemas", Salario: 6000.00}

    empresa.AdicionarEmpregado(emp1)
    empresa.AdicionarEmpregado(emp2)
    empresa.AdicionarEmpregado(emp3)

    empresa.ExibirEmpregados()
}
