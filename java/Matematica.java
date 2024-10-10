public class Matematica {
    public static int fatorial(int n) {
        if (n < 0) {
            throw new IllegalArgumentException("O fatorial não está definido para números negativos.");
        }
        if (n == 0 || n == 1) {
            return 1;
        }
        int resultado = 1;
        for (int i = 2; i <= n; i++) {
            resultado *= i;
        }
        return resultado;
    }

    public static void main(String[] args) {
        System.out.println(fatorial(5));  // Saída: 120
    }
}
