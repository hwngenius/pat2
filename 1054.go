package main

import "fmt"

func main() {
	var N, M, t int
	var m [1024*1024*16 + 1]int
	fmt.Scanf("%d %d", &M, &N)
	var sum = N*M/2 + 1
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scan(&t)
			m[t]++
			if m[t] >= sum {
				fmt.Println(t)
				return
			}
		}
	}
}

//20 18 out of time
