package main

import (
	"fmt"
)

type Escola struct {
	nome      string
	professores []*Professor
}

func (e *Escola) AdicionarProfessor(p *Professor) {
	for _, professor := range e.professores {
		if professor == p {
			return
		}
	}
	e.professores = append(e.professores, p)
	p.AdicionarEscola(e) 
}

func (e *Escola) RemoverProfessor(p *Professor) {
	for i, professor := range e.professores {
		if professor == p {
			e.professores = append(e.professores[:i], e.professores[i+1:]...) // Remove o professor
			p.RemoverEscola(e) 
			break
		}
	}
}

func (e Escola) String() string {
	return fmt.Sprintf("Escola: %s, Professores: %v", e.nome, e.professores)
}

type Professor struct {
	nome    string
	escolas []*Escola
}

func (p *Professor) AdicionarEscola(e *Escola) {
	for _, escola := range p.escolas {
		if escola == e {
			return 
		}
	}
	p.escolas = append(p.escolas, e) 
}

func (p *Professor) RemoverEscola(e *Escola) {
	for i, escola := range p.escolas {
		if escola == e {
			p.escolas = append(p.escolas[:i], p.escolas[i+1:]...) 
			break
		}
	}
}

func (p Professor) String() string {
	return fmt.Sprintf("Professor: %s, Escolas: %v", p.nome, p.escolas)
}


func main() {
	
	escolaA := &Escola{nome: "Escola A"}
	escolaB := &Escola{nome: "Escola B"}

	
	professorJoao := &Professor{nome: "Professor João"}
	professoraMaria := &Professor{nome: "Professora Maria"}


	escolaA.AdicionarProfessor(professorJoao)
	escolaA.AdicionarProfessor(professoraMaria)
	escolaB.AdicionarProfessor(professorJoao)


	fmt.Println(escolaA.String())
	fmt.Println(escolaB.String())
	fmt.Println(professorJoao.String())
	fmt.Println(professoraMaria.String())
}
