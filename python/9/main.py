from funcionario import *
from funcionario_horista import *
from funcionario_ass import *
if __name__ == "__main__":
    funcionario1 = FuncionarioHorista("Carlos", "123.456.789-00", 50.0, 160)
    funcionario2 = FuncionarioAssalariado("Ana", "987.654.321-00", 3000.0)

    print(f"Salário do Funcionario Horista: R$ {funcionario1.calcular_salario():.2f}")
    print(f"Salário do Funcionario Assalariado: R$ {funcionario2.calcular_salario():.2f}")
