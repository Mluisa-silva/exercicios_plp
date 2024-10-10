func main() {
	funcionario1 := FuncionarioHorista{"Carlos", "123.456.789-00", 50.0, 160}
	funcionario2 := FuncionarioAssalariado{"Ana", "987.654.321-00", 3000.0}

	fmt.Printf("Salário do Funcionario Horista: R$ %.2f\n", funcionario1.calcularSalario())
	fmt.Printf("Salário do Funcionario Assalariado: R$ %.2f\n", funcionario2.calcularSalario())
}
