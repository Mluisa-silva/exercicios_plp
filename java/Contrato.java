// Classe Contrato
public class Contrato implements Imprimivel {
    private String detalhes;

    public Contrato(String detalhes) {
        this.detalhes = detalhes;
    }

    @Override
    public void imprimir() {
        System.out.println("Imprimindo Contrato: " + detalhes);
    }
}
