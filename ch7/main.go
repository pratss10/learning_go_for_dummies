package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y, z float64
}

func (p point) radius() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	p := point{1, 2, 3}
	// t :=
	fmt.Println(p)
	p.x = 4
	fmt.Println(p)
	fmt.Println(p.radius())
}
