import java.util.ArrayList;
import java.util.List;

// Classe Professor
public class Professor {
    private String nome;
    private List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public String getNome() {
        return nome;
    }

    public List<Escola> getEscolas() {
        return escolas;
    }

    // Método para adicionar escola ao professor
    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this); // Associação bidirecional
        }
    }

    // Método para remover escola do professor
    public void removerEscola(Escola escola) {
        if (escolas.contains(escola)) {
            escolas.remove(escola);
            escola.removerProfessor(this); // Remoção bidirecional
        }
    }

    @Override
    public String toString() {
        return "Professor: " + nome;
    }
}
