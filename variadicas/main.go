package main

import "fmt"

func main() {
	fmt.Println(sum(1, 3, 4, 30, 200, 100))
}

func sum(numeros ...int) int {
	total := 0 
	for _, numero := range numeros {
		total += numero
	}
	return total
}