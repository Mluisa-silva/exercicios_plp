package main

import (
    "fmt"
    "errors"
)

type Matematica struct{}

func (m Matematica) Fatorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("o fatorial não está definido para números negativos")
    }
    if n == 0 || n == 1 {
        return 1, nil
    }
    resultado := 1
    for i := 2; i <= n; i++ {
        resultado *= i
    }
    return resultado, nil
}

func main() {
    m := Matematica{}
    resultado, err := m.Fatorial(5)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(resultado)
    }
}