package main

import "fmt"

func double(s string) string {
	var b []byte
	var t int
	for i := len(s) - 1; i >= 0; i-- {
		tmp := int(s[i]-'0')*2%10 + t
		b = append(b, byte(tmp+'0'))
		t = int(s[i]-'0') * 2 / 10
	}
	if byte(t+'0') != '0' {
		b = append(b, byte(t+'0'))
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	var s string = "5"
	m := make(map[byte]int)
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	s = double(s)
	for i := 0; i < len(s); i++ {
		m[s[i]]--
	}
	mask := 0
	for _, v := range m {
		if v != 0 {
			mask = 1
			break
		}
	}
	if mask == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	fmt.Println(s)
}
