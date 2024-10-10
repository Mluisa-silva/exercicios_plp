class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0

    def acelerar(self, incremento):
        if incremento > 0:
            self.velocidade += incremento
            print(f"O carro acelerou {incremento} km/h.")
        else:
            print("O valor de aceleração deve ser positivo.")

    def frear(self, decremento):
        if decremento > 0 and self.velocidade - decremento >= 0:
            self.velocidade -= decremento
            print(f"O carro freou {decremento} km/h.")
        else:
            print("A frenagem não pode resultar em velocidade negativa.")

    def exibir_velocidade(self):
        print(f"Velocidade atual: {self.velocidade} km/h")

    def exibir_detalhes(self):
        print(f"Marca: {self.marca}")
        print(f"Modelo: {self.modelo}")
        print(f"Ano: {self.ano}")
