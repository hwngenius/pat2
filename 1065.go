package main

import (
	"fmt"
)

func main() {
	var N int
	var a, b, c int64
	var flag bool
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d %d", &a, &b, &c)
		fmt.Printf("Case #%d: ", i+1)
		res := a + b
		if a > 0 && b > 0 && res < 0 {
			flag = true
		} else if a < 0 && b < 0 && res >= 0 {
			flag = false
		} else if res > c {
			flag = true
		} else {
			flag = false
		}
		if flag {
			fmt.Printf("true\n")
		} else {
			fmt.Printf("false\n")
		}
	}
}
