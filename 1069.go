package main

import (
	"fmt"
	"sort"
	"strconv"
)

type mySlice []byte

func (m mySlice) Len() int           { return len(m) }
func (m mySlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m mySlice) Less(i, j int) bool { return m[i] > m[j] }
func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	t := 0
	for t != 6174 {
		var b []byte
		for i := 0; i < len(s); i++ {
			b = append(b, s[i])
		}
		for len(b) != 4 {
			b = append(b, '0')
		}
		sort.Sort(mySlice(b))
		fmt.Printf("%s - ", string(b))
		t1, _ := strconv.Atoi(string(b))
		reverse(b)
		fmt.Printf("%s = ", string(b))
		t2, _ := strconv.Atoi(string(b))
		t = t1 - t2
		fmt.Printf("%04d\n", t)
		if t == 0 {
			return
		}
		s = strconv.Itoa(t)
	}

}
