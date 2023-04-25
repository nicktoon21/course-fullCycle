package main

type MyNumber int 

//Constraints
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string] T) T{
	var soma T
	for _, value := range m {
		soma += value
	}
	return soma
}

func Compara[T comparable](a , b T) bool{
	if a == b {
		return true
	}
	return false
}

func main(){
	m := map[string] int{
		"Nicolas":1000,
		"Leonardo":3000,
		"Italo":2000,
	}
	m1 := map[string] MyNumber{
		"Nicolas":1000,
		"Leonardo":3000,
		"Italo":2000,
	}
	f := map[string] float64{
		"Nicolas":1000.0,
		"Leonardo":3000.1,
		"Italo":2000.2,
	}

	println(Soma(m))
	println(Soma(m1))
	println(Soma(f))
	println(Compara(10 , 10))

}