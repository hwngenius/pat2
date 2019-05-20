package main

import "fmt"

func main() {
	var sum float64 = 1
	var t, max float64
	var index int
	var set = [...]byte{'W', 'T', 'L'}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&t)
			if t > max {
				max = t
				index = j
			}
		}
		sum *= max
		fmt.Printf("%c ", set[index])
		max = 0
	}
	fmt.Printf("%.2f", (sum*0.65-1)*2)
}
