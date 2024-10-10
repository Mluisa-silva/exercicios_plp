func main() {
    produto1 := Produto{"Produto A", 10.00}
    produto2 := Produto{"Produto B", 15.50}
    produtoSoma := produto1.Soma(produto2)
    fmt.Println(produtoSoma)
}