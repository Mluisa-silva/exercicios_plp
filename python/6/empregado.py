class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def exibir_detalhes(self):
        print(f"Nome: {self.nome}")
        print(f"Cargo: {self.cargo}")
        print(f"Salário: R$ {self.salario:.2f}")

class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def exibir_empregados(self):
        print(f"Empresa: {self.nome}")
        print("Lista de Empregados:")
        for empregado in self.empregados:
            empregado.exibir_detalhes()
            print("---------------")

