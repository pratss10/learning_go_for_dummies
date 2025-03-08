package main

import (
	"fmt"
)

func main() {
	var cars [3]string
	cars[0] = "Porsche"
	cars[1] = "Rolls-Royce"
	cars[2] = "Bugatti"

	for i, v := range cars {
		fmt.Println(i, v)
	}

	for i, v := range "Pratyaksh Sharma" {
		fmt.Printf(("%d %c\n"), i, v)
	}
}
