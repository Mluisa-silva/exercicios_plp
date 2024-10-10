from saldo import *
if __name__ == "__main__":
    conta = ContaBancaria(100.0)
    try:
        conta.sacar(150.0)
    except SaldoInsuficienteException as e:
        print(e)