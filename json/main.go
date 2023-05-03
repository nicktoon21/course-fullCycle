package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo int `json:"s"`
}

func main() {
	conta := Conta{Numero:1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}


	jsonPuro := []byte(`{"n":2, "s":200}`)
	var contax Conta
	err = json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		panic(err)
	}
	println(contax.Saldo)
}