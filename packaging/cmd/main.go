package main

import (
	"fmt"

	"github.com/nicktoon21/packaging/math"
)

func main() {

	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
