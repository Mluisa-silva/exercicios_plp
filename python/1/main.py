from Carro import *

if __name__ == "__main__":
    carro1 = Carro("Toyota", "Corolla", 2020)

    carro1.exibir_detalhes()

    carro1.acelerar(30)
    carro1.exibir_velocidade()

    carro1.acelerar(20)
    carro1.exibir_velocidade()

    carro1.frear(10)
    carro1.exibir_velocidade()

    carro1.frear(50)
    carro1.exibir_velocidade()
