package main

import "fmt"


func soma(a, b *int) int {
	return *a + *b;
}
func main() {
	ponteiro1 := 10
	ponteiro2 := 20
	a := soma(&ponteiro1, &ponteiro2)
	fmt.Println(a)
}