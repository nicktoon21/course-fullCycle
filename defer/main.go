package main

import "fmt"
func main() {
	defer fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}