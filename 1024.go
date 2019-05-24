package main

import (
	"fmt"
)

func isPalindromic(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	var r []byte
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return string(r)
}

func btoi(b byte) int {
	return int(b - '0')
}
func itob(i int) byte {
	return byte(i + '0')
}

func add(s1, s2 string) string {
	var r []byte
	var jin = 0
	for i, j := len(s1)-1, len(s2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		t := btoi(s1[i]) + btoi(s2[j])
		b := itob((t + jin) % 10)
		r = append(r, b)
		jin = t / 10
	}
	if jin != 0 {
		r = append(r, itob(jin))
	}
	return reverse(string(r))
}

func main() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if isPalindromic(s) {
			fmt.Println(s)
			fmt.Println(i)
			return
		}
		s = add(s, reverse(s))
	}
	fmt.Println(s)
	fmt.Println(n)
}

// 25/18 wrong answer
