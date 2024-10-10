import java.util.ArrayList;
import java.util.List;

public class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public String getNome() {
        return nome;
    }

    public List<Professor> getProfessores() {
        return professores;
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
            professor.adicionarEscola(this); // Associação bidirecional
        }
    }

    public void removerProfessor(Professor professor) {
        if (professores.contains(professor)) {
            professores.remove(professor);
            professor.removerEscola(this); // Remoção bidirecional
        }
    }

    @Override
    public String toString() {
        return "Escola: " + nome;
    }
}

