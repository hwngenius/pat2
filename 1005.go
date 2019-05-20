package main

import "fmt"

func f(n int) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return ""
	}
}

func main() {
	var s string
	var sum int
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
	}
	if sum == 0 {
		fmt.Println("zero")
		return
	}
	i := 0
	for sum > 0 {
		if i > 0 {
			defer fmt.Printf("%s ", f(sum%10))
		} else {
			defer fmt.Printf("%s", f(sum%10))
			i = 1
		}
		sum /= 10
	}
}
