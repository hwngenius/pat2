package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
		t := unicode.ToUpper(rune(s2[i]))
		set[int(t)] = true
	}
	for i := 0; i < len(s1); i++ {
		t := unicode.ToUpper(rune(s1[i]))
		if set[int(t)] == false {
			fmt.Printf("%c", t)
			set[int(t)] = true
		}
	}
}
