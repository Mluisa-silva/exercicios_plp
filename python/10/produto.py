class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, other):
        if isinstance(other, Produto):
            return Produto(f"{self.nome} + {other.nome}", self.preco + other.preco)
        return NotImplemented

    def __repr__(self):
        return f"{self.nome}: R${self.preco:.2f}"


produto1 = Produto("Produto A", 10.00)
produto2 = Produto("Produto B", 15.50)
produto_soma = produto1 + produto2
print(produto_soma) 
