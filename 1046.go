package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var N int
	var sum int
	var set [100000]int
	var dis [100000]int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&set[i])
		sum += set[i]
		dis[i] = sum
	}
	var M, a, b int
	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d", &a, &b)
		t := 0
		if a > b {
			a, b = b, a
		}
		if a == 1 {
			t = dis[b-2]
		} else {
			t = dis[b-2] - dis[a-2]
		}
		fmt.Println(min(t, sum-t))
	}
}

// 20 17 out of time
