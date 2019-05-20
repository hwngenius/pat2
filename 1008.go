package main

import "fmt"

func main() {
	var N int
	var sum, current, t int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&t)
		sum += 5
		if t > current {
			sum += (t - current) * 6
		} else {
			sum += (current - t) * 4
		}
		current = t
		//fmt.Println(sum)
	}
	fmt.Println(sum)
}
