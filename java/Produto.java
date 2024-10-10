class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public Produto soma(Produto outro) {
        return new Produto(this.nome + " + " + outro.nome, this.preco + outro.preco);
    }

    @Override
    public String toString() {
        return nome + ": R$" + String.format("%.2f", preco);
    }

    public static void main(String[] args) {
        Produto produto1 = new Produto("Produto A", 10.00);
        Produto produto2 = new Produto("Produto B", 15.50);
        Produto produtoSoma = produto1.soma(produto2);
        System.out.println(produtoSoma);  // Saída: Produto A + Produto B: R$25.50
    }
}
