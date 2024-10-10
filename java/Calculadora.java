public class Calculadora {

    public double somar(double a, double b) {
        return a + b;
    }

    public double somar(double a, double b, double c) {
        return a + b + c;
    }

    public static void main(String[] args) {
        Calculadora calc = new Calculadora();

        double resultado1 = calc.somar(5.0, 3.0);
        System.out.println("Soma de 5.0 e 3.0: " + resultado1);

        // Testando a soma de três números
        double resultado2 = calc.somar(5.0, 3.0, 2.0);
        System.out.println("Soma de 5.0, 3.0 e 2.0: " + resultado2);
    }
}
