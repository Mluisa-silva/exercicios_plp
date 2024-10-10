public class Carro1 {
    public static void main(String[] args) {
        Carro carro1 = new Carro("Toyota", "Corolla", 2020);

        carro1.exibirDetalhes();

        carro1.acelerar(30);
        carro1.exibirVelocidade();

        carro1.acelerar(20);
        carro1.exibirVelocidade();

        carro1.frear(10);
        carro1.exibirVelocidade();

        carro1.frear(50);
        carro1.exibirVelocidade();
    }
}
