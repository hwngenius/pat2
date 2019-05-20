package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var s string
	for i := 0; i < 2; i++ {
		fmt.Scan(&s)
		L := strings.Split(s, ".")
		if t, e := strconv.Atoi(L[0]); e == nil {
			sum += t * 17 * 29
		} else {
			return
		}
		if t, e := strconv.Atoi(L[1]); e == nil {
			sum += t * 29
		} else {
			return
		}
		if t, e := strconv.Atoi(L[2]); e == nil {
			sum += t
		} else {
			return
		}
	}
	fmt.Printf("%d.%d.%d", sum/(29*17), sum%(29*17)/29, sum%29)
}
