package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int = 1
	var s string
	var m1 = [...]string{"jan", "feb", "mar", "apr", "may", "jun", "jly", "aug", "sep", "oct", "nov", "dec"}
	var m2 = [...]string{"tam", "hel", "maa", "huh", "tou", "kes", "hei", "elo", "syy", "lok", "mer", "jou"}
	fmt.Scan(&N)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s = input.Text()
		t, err := strconv.Atoi(s)
		if err != nil {
			L := strings.Split(s, " ")
			n := 0
			for i := 0; i < len(L); i++ {
				j := 0
				for j = 0; j < len(m1); j++ {
					if m1[j] == L[i] {
						n += j + 1
					}
				}
				for j = 0; j < len(m2); j++ {
					if m2[j] == L[i] {
						n += (j + 1) * 13
					}
				}
			}
			fmt.Println(n)
		} else {
			if t >= 13 {
				t1 := t / 13
				t %= 13
				fmt.Printf("%s", m2[t1-1])
				if t != 0 {
					fmt.Printf(" %s\n", m1[t-1])
				} else {
					fmt.Printf("\n")
				}
			} else {
				if t != 0 {
					fmt.Printf("%s\n", m1[t-1])
				} else {
					fmt.Printf("tret\n")
				}
			}
		}
	}

}
