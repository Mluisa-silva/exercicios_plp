
class SaldoInsuficienteException extends Exception {
    public SaldoInsuficienteException(String mensagem) {
        super(mensagem);
    }
}

class ContaBancaria2 {
    private double saldo;

    public ContaBancaria2(double saldoInicial) {
        this.saldo = saldoInicial;
    }

    public void sacar(double valor) throws SaldoInsuficienteException {
        if (valor > saldo) {
            throw new SaldoInsuficienteException("Saldo insuficiente para realizar o saque.");
        }
        saldo -= valor;
    }

    public double getSaldo() {
        return saldo;
    }
}

public class SaldoMain {
    public static void main(String[] args) {
        ContaBancaria2 conta = new ContaBancaria2(100.0);
        try {
            conta.sacar(150.0);
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }
}
