package main

func main() {
	a := 1
	b := 2
	c := 3
	d := 0
	if a > b {
		println(a)
	} else {
		println(b)
	}
	if a == b {
		println(a)
	} else {
		println(b)
	}
	if a < b && c > a {
		println(a)
	} else {
		println(b)
	}
	if a < b || c > a {
		println(a)
	} else {
		println(b)
	}

	switch d {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("null")
	}
}
