class Motor:
    def __init__(self, potencia, tipo):
        self.potencia = potencia
        self.tipo = tipo

    def get_potencia(self):
        return self.potencia

    def set_potencia(self, potencia):
        self.potencia = potencia

    def get_tipo(self):
        return self.tipo

    def set_tipo(self, tipo):
        self.tipo = tipo

    def exibir_informacoes(self):
        print(f"Potência do motor: {self.potencia} CV")
        print(f"Tipo de motor: {self.tipo}")
