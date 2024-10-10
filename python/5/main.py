from escola import *

if __name__ == "__main__":
    escola_a = Escola("Escola A")
    escola_b = Escola("Escola B")

    professor_joao = Professor("Professor João")
    professora_maria = Professor("Professora Maria")

    escola_a.adicionar_professor(professor_joao)
    escola_a.adicionar_professor(professora_maria)

    escola_b.adicionar_professor(professor_joao)

    print(escola_a.nome + " tem os professores: " + ", ".join([professor.nome for professor in escola_a.professores]))
    print(escola_b.nome + " tem os professores: " + ", ".join([professor.nome for professor in escola_b.professores]))

    print(professor_joao.nome + " leciona nas escolas: " + ", ".join([escola.nome for escola in professor_joao.escolas]))
    print(professora_maria.nome + " leciona nas escolas: " + ", ".join([escola.nome for escola in professora_maria.escolas]))