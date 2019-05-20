package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s1, s2 string
	var set int
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Text()
	input.Scan()
	s1 = input.Text()
	input.Scan()
	s2 = input.Text()
	for i, j := len(s1)-1, len(s2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if s1[i] == s2[j] {
			set++
		} else {
			break
		}
	}
	if set == 0 {
		fmt.Println("nail")
		return
	}
	for input.Scan() {
		s2 = input.Text()
		for s1[len(s1)-set] != s2[len(s2)-set] {
			set--
		}
		if set == 0 {
			fmt.Println("nail")
			return
		}
	}
	fmt.Println(s1[len(s1)-set:])
}
