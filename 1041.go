package main

import "fmt"

func main() {
	var N int
	var ans [100001]int
	var set [10001]int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ans[i])
		set[ans[i]]++
	}
	a := -1
	for i := 0; i < N; i++ {
		if set[ans[i]] == 1 {
			a = ans[i]
			break
		}
	}
	if a == -1 {
		fmt.Println("None")
	} else {
		fmt.Println(a)
	}
}

//16 20 out of time
