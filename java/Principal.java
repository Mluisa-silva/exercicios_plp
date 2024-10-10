public class Principal {
    public static void main(String[] args) {
        Imprimivel relatorio = new Relatório("Relatório Financeiro");
        Imprimivel contrato = new Contrato("Contrato de Prestação de Serviços");

        relatorio.imprimir();
        contrato.imprimir();
    }
}
