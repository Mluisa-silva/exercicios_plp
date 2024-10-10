package main

import (
    "sync"
)

type Configuracao struct {
    config map[string]string
}

var (
    instance *Configuracao
    once     sync.Once
)

func GetInstance() *Configuracao {
    once.Do(func() {
        instance = &Configuracao{
            config: make(map[string]string),
        }
    })
    return instance
}

func (c *Configuracao) SetConfig(chave, valor string) {
    c.config[chave] = valor
}

func (c *Configuracao) GetConfig(chave string) string {
    return c.config[chave]
}

func main() {
    config1 := GetInstance()
    config1.SetConfig("tema", "escuro")

    config2 := GetInstance()
    println(config2.GetConfig("tema"))
    println(config1 == config2)          
}
