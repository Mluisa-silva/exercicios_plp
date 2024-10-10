package main

import (
    "fmt"
)

func somarDois(a float64, b float64) float64 {
    return a + b
}

func somarTres(a float64, b float64, c float64) float64 {
    return a + b + c
}

func main() {
    resultado1 := somarDois(5.0, 3.0)
    fmt.Println("Soma de 5.0 e 3.0:", resultado1)
	
    resultado2 := somarTres(5.0, 3.0, 2.0)
    fmt.Println("Soma de 5.0, 3.0 e 2.0:", resultado2)
}
