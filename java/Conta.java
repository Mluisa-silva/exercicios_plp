public class Conta{
    public static void main(String[] args) {
        ContaBancaria conta = new ContaBancaria("Luísa Silva", 1000.0);

        System.out.println("Titular: " + conta.getTitular());
        System.out.println("Saldo inicial: R$ " + conta.getSaldo());

        conta.depositar(1000.0);
        System.out.println("Saldo: R$ " + conta.getSaldo());

        conta.sacar(290.0);
        System.out.println("Saldo: R$ " + conta.getSaldo());
    }
}
