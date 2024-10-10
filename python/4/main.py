from motor import *
from carro import *

if __name__ == "__main__":
    motor1 = Motor(150, "Gasolina")

    carro1 = Carro("Toyota", "Corolla", motor1)

    carro1.exibir_informacoes()
