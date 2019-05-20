package main

import "fmt"

func main() {
	var s string = "helloworld!"
	fmt.Scan(&s)
	var n1 = (len(s) + 2) / 3
	var n2 = len(s) + 2 - 2*n1
	for i := 0; i < n1-1; i++ {
		for j := 0; j < n2; j++ {
			if j == 0 {
				fmt.Printf("%c", s[i])
			} else if j == n2-1 {
				fmt.Printf("%c", s[len(s)-1-i])
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	for i := 0; i < n2; i++ {
		fmt.Printf("%c", s[i+n1-1])
	}
}
