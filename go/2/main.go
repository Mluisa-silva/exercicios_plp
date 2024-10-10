package main

func main() {
    conta := NovoContaBancaria("João Silva", 1000.0)

    println("Titular:", conta.GetTitular())
    println("Saldo inicial: R$", conta.GetSaldo())

    conta.Depositar(500.0)
    println("Saldo após depósito: R$", conta.GetSaldo())

    conta.Sacar(300.0)
    println("Saldo após saque: R$", conta.GetSaldo())
}
