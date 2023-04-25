package main

import (
	"errors"
	"fmt"
)

const textoError = "a soma e maior que 50"

func main() {
	valor, err := sum(1, 49)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(valor)
	}
}

func sum(a ,b int) (int, error) {
	if a + b >= 50 {
		return a + b, errors.New(textoError)
	}
	return a + b, nil
}