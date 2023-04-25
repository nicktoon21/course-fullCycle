package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	
	a := append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
	
}
