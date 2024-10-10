// Classe Relatório
public class Relatório implements Imprimivel {
    private String conteudo;

    public Relatório(String conteudo) {
        this.conteudo = conteudo;
    }

    @Override
    public void imprimir() {
        System.out.println("Imprimindo Relatório: " + conteudo);
    }
}
