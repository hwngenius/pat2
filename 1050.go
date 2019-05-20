package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	var set [256]bool
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s1 = input.Text()
	input.Scan()
	s2 = input.Text()
	for i := 0; i < len(s2); i++ {
		set[int(s2[i])] = true
	}
	for i := 0; i < len(s1); i++ {
		if !set[int(s1[i])] {
			fmt.Printf("%c", s1[i])
		}
	}
}
