package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo float64
}

func (c Conta) simular(valor float64) float64 {
	 c.saldo += valor
	 return c.saldo
}

func (c Cliente) andou(){
	c.nome = "Naldo Ferreira"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func NewConta() *Conta {
	return &Conta{saldo:0}
}

func main() {
	naldo := Cliente{
		nome: "Naldo",
	}
	naldo.andou()
	fmt.Printf("O nome que está na minha struct é %v\n", naldo.nome)

	conta := Conta{saldo: 600}
	simulacao := conta.simular(200)
	fmt.Printf("O valor simulado é R$ %v\n", simulacao)
	fmt.Printf("O saldo da conta é R$ %v\n",conta.saldo)

	c:= NewConta()
	fmt.Printf("Saldo da nova conta R$ %v\n",c.saldo)
}