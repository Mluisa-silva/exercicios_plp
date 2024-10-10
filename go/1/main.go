// main.go
package main

func main() {
	carro1 := Carro{marca: "Toyota", modelo: "Corolla", ano: 2020}

	carro1.ExibirDetalhes()
	
	carro1.Acelerar(30)
	carro1.ExibirVelocidade()

	carro1.Acelerar(20)
	carro1.ExibirVelocidade()

	carro1.Frear(10)
	carro1.ExibirVelocidade()

	carro1.Frear(50)
	carro1.ExibirVelocidade()
}
