package main

import (
	"fmt"
	"math"
)

func isPrimer(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	t := int(math.Sqrt(float64(n))) + 1
	for i := 2; i <= t; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var N, D int

	for {
		fmt.Scan(&N)
		// fmt.Printf("%d ", N)
		if N < 0 {
			break
		}
		fmt.Scan(&D)
		if isPrimer(N) == false {
			fmt.Println("No")
			continue
		}
		var i int
		var t int = D
		for t <= N {
			t *= D
			i++
		}
		var m []int
		t /= D
		for i >= 0 {
			m = append(m, N/t)
			N %= t
			t /= D
			i--
		}
		t = 1
		sum := 0
		for i := 0; i < len(m); i++ {
			sum += m[i] * t
			t *= D
		}
		if isPrimer(sum) == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
