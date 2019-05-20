package main

import "fmt"

func isRgitht(set []int) bool {
	for i, j := 0, len(set)-1; i < j; i, j = i+1, j-1 {
		if set[i] != set[j] {
			return false
		}
	}
	return true
}

func main() {
	var N, b int = 27, 2
	var set []int
	fmt.Scanf("%d %d", &N, &b)
	// if b > N {
	// 	fmt.Println("Yes")
	// 	fmt.Println(N)
	// 	return
	// }
	var i int
	var t int = b
	for t <= N {
		t *= b
		i++
	}
	t /= b
	for i >= 0 {
		set = append(set, N/t)
		N %= t
		t /= b
		i--
	}
	if isRgitht(set) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	for i2 := 0; i2 < len(set); i2++ {
		if i2 != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", set[i2])
	}
}
