class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)
            professor.adicionar_escola(self)

    def remover_professor(self, professor):
        if professor in self.professores:
            self.professores.remove(professor)
            professor.remover_escola(self) 
    def __str__(self):
        return f"Escola: {self.nome}"


class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)

    def remover_escola(self, escola):
        if escola in self.escolas:
            self.escolas.remove(escola)
            escola.remover_professor(self)

    def __str__(self):
        return f"Professor: {self.nome}"

