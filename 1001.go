package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	//fmt.Println(a + b)
	s := strconv.Itoa(a + b)
	mask := 0
	if s[0] == '-' {
		mask = 1
	} else {
		mask = 0
	}
	for j, i := 1, len(s)-1; i >= 0; i, j = i-1, j+1 {
		if j%3 == 0 && i > mask {
			defer fmt.Printf(",%c", s[i])
		} else {
			defer fmt.Printf("%c", s[i])
		}
	}
}
