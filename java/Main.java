public class Main {
    public static void main(String[] args) {
        Escola escola1 = new Escola("Escola A");
        Escola escola2 = new Escola("Escola B");
        
        Professor professor1 = new Professor("Professor João");
        Professor professor2 = new Professor("Professora Maria");
        
        escola1.adicionarProfessor(professor1);
        escola1.adicionarProfessor(professor2);

        escola2.adicionarProfessor(professor1);
        
        System.out.println(escola1.getNome() + " tem os professores: " + escola1.getProfessores());
        System.out.println(escola2.getNome() + " tem os professores: " + escola2.getProfessores());

        System.out.println(professor1.getNome() + " leciona nas escolas: " + professor1.getEscolas());
        System.out.println(professor2.getNome() + " leciona nas escolas: " + professor2.getEscolas());
    }
}
