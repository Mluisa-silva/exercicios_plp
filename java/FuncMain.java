public class FuncMain {
    public static void main(String[] args) {
        Funcionario funcionario1 = new FuncionarioHorista("Carlos", "123.456.789-00", 50.0, 160);
        Funcionario funcionario2 = new FuncionarioAssalariado("Ana", "987.654.321-00", 3000.0);

        System.out.println("Salário do Funcionario Horista: R$ " + funcionario1.calcularSalario());
        System.out.println("Salário do Funcionario Assalariado: R$ " + funcionario2.calcularSalario());
    }
}
