package main

import "fmt"

func main() {
	var N int
	var letter = [...]byte{'S', 'H', 'C', 'D', 'J'}
	fmt.Scan(&N)
	var order [54]int
	for i := 0; i < 54; i++ {
		fmt.Scan(&order[i])
	}
	var ans [54]int
	for i := 0; i < 54; i++ {
		ans[i] = i
	}
	var tmp [54]int
	for j := 0; j < N; j++ {
		for i := 0; i < 54; i++ {
			tmp[order[i]-1] = ans[i]
		}
		ans = tmp
	}
	for i := 0; i < 54; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%c%d", letter[ans[i]/13], ans[i]%13+1)
	}
}
