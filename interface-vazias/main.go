package main

import "fmt"

func main() {
	// var x interface {} = 10 
	// var y interface {} = "Olá mundo!"

	// showType(x)
	// showType(y)

	var minhaVar interface {} = "Nicolas Aguiar"
	fmt.Println(minhaVar.(string))

	res, ok := minhaVar.(int)
	if ok {
		fmt.Println(res)
	}
	fmt.Println("Minha variável não foi convertida")
}


// func showType(t interface{}) {
// 	fmt.Printf("O tipo da variavel é %T o valor é %v\n ", t, t )
// }