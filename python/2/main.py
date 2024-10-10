from conta_bancaria import *

def main():

    conta = ContaBancaria("João Silva", 1000.0)

    # Exibindo o titular e o saldo inicial
    print("Titular:", conta.get_titular())
    print("Saldo inicial: R$ {:.2f}".format(conta.get_saldo()))

    # Realizando depósito e saque
    conta.depositar(500.0)
    print("Saldo após depósito: R$ {:.2f}".format(conta.get_saldo()))

    conta.sacar(300.0)
    print("Saldo após saque: R$ {:.2f}".format(conta.get_saldo()))

if __name__ == "__main__":
    main()
