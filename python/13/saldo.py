
class SaldoInsuficienteException(Exception):
    def __init__(self, mensagem):
        super().__init__(mensagem)

class ContaBancaria:
    def __init__(self, saldo_inicial):
        self.saldo = saldo_inicial

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException("Saldo insuficiente para realizar o saque.")
        self.saldo -= valor

    def get_saldo(self):
        return self.saldo

