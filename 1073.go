package main

import (
	"fmt"
	"strconv"
)

func prechange(s string) (t string) {
	l := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '.' {
			l = append(l, s[i])
		}
	}
	t = string(l)
	return
}

func bacChange(s string, b int) (r string) {
	l := []byte{}
	mask := 0
	for i := 0; i < len(s); i++ {
		if mask == 1 {
			b--
			if b == -1 {
				l = append(l, '.')
			}
			l = append(l, s[i])
		} else {
			if s[i] != '.' {
				l = append(l, s[i])
			} else {
				mask = 1
			}
		}
	}
	r = string(l)
	for ; b > 0; b-- {
		r += "0"
	}
	return
}

func main() {
	var a = "-1.2E-1"
	var b, pre string
	var i int
	fmt.Scanf("%s", &a)
	for i = 0; i < len(a); i++ {
		if a[i] == 'E' {
			break
		}
	}
	a, b = a[:i], a[i+1:]
	if a[0] == '-' {
		fmt.Printf("%c", a[0])
	}
	a = a[1:]
	//fmt.Println(a, b)
	i, _ = strconv.Atoi(b)
	if i == 0 {
		fmt.Println(a)
		return
	} else if i < 0 {
		pre = "0."
		for i < -1 {
			pre += "0"
			i++
		}
		a = prechange(a)
		a = pre + a
		fmt.Println(a)
	} else {
		a = bacChange(a, i)
		fmt.Println(a)
	}
}
