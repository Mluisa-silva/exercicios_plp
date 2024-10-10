public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    public void acelerar(int incremento) {
        if (incremento > 0) {
            this.velocidade += incremento;
            System.out.println("O carro acelerou " + incremento + " km/h.");
        } else {
            System.out.println("Digite um valor válido.");
        }
    }

    public void frear(int decremento) {
        if (decremento > 0 && this.velocidade - decremento >= 0) {
            this.velocidade -= decremento;
            System.out.println("O carro freou " + decremento + " km/h.");
        } else {
            System.out.println("Valor inválido.");
        }
    }

    public void exibirVelocidade() {
        System.out.println("Velocidade atual: " + this.velocidade + " km/h.");
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + this.marca);
        System.out.println("Modelo: " + this.modelo);
        System.out.println("Ano: " + this.ano);
    }
}
