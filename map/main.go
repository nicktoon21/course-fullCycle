package main

import "fmt"

func main() {
	salarios:= map[string]int{
		"Nando":1000, 
		"João": 2000,
		"Mariana": 3000,
	}
	// fmt.Println(salarios["Mariana"])
	// delete(salarios, "Mariana")
	// salarios["Mariana"] = 3500;
	// fmt.Println(salarios["Mariana"])


	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal1["Joaquim"] = 1000
	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é R$ %d\n", nome, salario)
	}
}