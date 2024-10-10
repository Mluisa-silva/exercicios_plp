package main

import "fmt"

type Animal interface {
	Som()
}

func EmitirSomDosAnimais(animais []Animal) {
	for _, animal := range animais {
		animal.Som() 
	}
}

