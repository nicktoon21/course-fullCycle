package main

import "fmt"

type Endereco struct {
	Lagradouro string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
	Endereco //Address Endereco
}

type Funcionario struct {
	Nome string
}


//interface só pode passar metodos, não podemos implementar propriedades
type Pessoa interface {
	Desativar()
}

func (f Funcionario) Desativar(){

}
func (c *Cliente) Desativar(){
	c.Ativo = false
	fmt.Printf("O cliente %v foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa){
	pessoa.Desativar()
}

func main() {
	naldo := Cliente{
		Nome: "Naldo",
		Idade: 37,
		Ativo: true,
	}

	marioF:= Funcionario{}

	Desativacao(marioF)//naldo.Desativar()
	Desativacao(&naldo)//naldo.Desativar()
	naldo.Cidade = "Maranhão" // naldo.Address.Cidade = "Maranhão"

	fmt.Println(naldo.Cidade) // fmt.Println(naldo.Address.Cidade) 
}