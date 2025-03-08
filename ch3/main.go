package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	if num&1 == 1 {
		fmt.Println(num, "is odd")
	} else {
		fmt.Println(num, "is even")
	}
}
