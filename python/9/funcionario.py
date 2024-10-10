from abc import ABC, abstractmethod

class Funcionario(ABC):
    def __init__(self, nome, cpf):
        self.nome = nome
        self.cpf = cpf

    @abstractmethod
    def calcular_salario(self):
        pass
