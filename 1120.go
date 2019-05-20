package main

import (
	"fmt"
	"sort"
)

func Id(s string) (n int) {
	for i := 0; i < len(s); i++ {
		n += int(s[i] - '0')
	}
	return n
}

type myslice []int

func (m myslice) Len() int {
	return len(m)
}
func (m myslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m myslice) Less(i, j int) bool {
	return m[i] < m[j]
}
func main() {
	var set [10000]bool
	var output []int
	var N int
	var s string
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&s)
		t := Id(s)
		if set[t] == false {
			output = append(output, t)
			set[t] = true
		}
	}
	sort.Sort(myslice(output))
	fmt.Println(len(output))
	for i := 0; i < len(output); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", output[i])
	}
}
