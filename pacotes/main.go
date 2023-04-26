package main

import (
	"fmt"

	c "github.com/nicktoon21/course-fullCycle/pacotes/carro"
	m "github.com/nicktoon21/course-fullCycle/pacotes/matematica"
)

func main() {
	soma := m.Soma(1, 5)
	carro := c.Carro{}
	car := &soma
	fmt.Printf("Resultado %v\n", car)

	carro.SetKm(0.0)
	carro.SetMarca("vw")
	carro.SetModelo("Polo")

	fmt.Printf("Resultado %v\n", soma)
	fmt.Printf("Resultado %v\n", carro.GetKm())
	fmt.Printf("Resultado %v\n", carro.GetMarca())
	fmt.Printf("Resultado %v\n", carro.GetModelo())
}
