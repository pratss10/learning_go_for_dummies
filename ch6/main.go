package main

import "fmt"

func main() {
	var tc int
	fmt.Scan(&tc)
	for range tc {
		var s string
		fmt.Scan(&s)
		n := len(s)
		var ans int = n
		for i := n - 2; i >= 0; i-- {
			if s[i] == s[i+1] {
				ans = 1
				break
			}
		}
		fmt.Println(ans)
	}
}
