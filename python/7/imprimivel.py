from abc import ABC, abstractmethod

class Imprimível(ABC):
    
    @abstractmethod
    def imprimir(self):
        pass

class Relatório(Imprimível):
    def __init__(self, conteudo):
        self.conteudo = conteudo

    def imprimir(self):
        print(f"Imprimindo Relatório: {self.conteudo}")

class Contrato(Imprimível):
    def __init__(self, detalhes):
        self.detalhes = detalhes

    def imprimir(self):
        print(f"Imprimindo Contrato: {self.detalhes}")

