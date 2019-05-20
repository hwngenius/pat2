package main

import (
	"fmt"
	"strconv"
	"strings"
)

func change(s string) (t int) {
	L := strings.Split(s, ":")
	t1, _ := strconv.Atoi(L[0])
	t += t1 * 3600
	t1, _ = strconv.Atoi(L[1])
	t += t1 * 60
	t1, _ = strconv.Atoi(L[2])
	t += t1
	if t == 24*60*60 {
		t = 0
	}
	return t
}

func main() {
	var first, last string
	var max, min int
	max = 0
	min = 24*60*60 - 1
	var N int
	var id, signin, signout string
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %s %s", &id, &signin, &signout)
		if change(signin) < min {
			min = change(signin)
			first = id
		}
		if change(signout) > max {
			max = change(signout)
			last = id
		}
	}
	fmt.Println(first, last)
}
