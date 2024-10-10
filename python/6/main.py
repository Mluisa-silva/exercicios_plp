from empregado import *
if __name__ == "__main__":
    empresa = Empresa("Tech Solutions")


    emp1 = Empregado("João Silva", "Desenvolvedor", 5000.00)
    emp2 = Empregado("Maria Oliveira", "Gerente de Projetos", 8000.00)
    emp3 = Empregado("Pedro Santos", "Analista de Sistemas", 6000.00)

    empresa.adicionar_empregado(emp1)
    empresa.adicionar_empregado(emp2)
    empresa.adicionar_empregado(emp3)


    empresa.exibir_empregados()