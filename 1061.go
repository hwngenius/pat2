package main

import (
	"fmt"
	"unicode"
)

func main() {
	var WEEK = [...]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	var s1, s2, s3, s4 string
	var mask bool
	var week, hour, minute int
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Scan(&s3)
	fmt.Scan(&s4)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] == s2[i] {
			if mask {
				if s1[i] >= 'A' && s1[i] <= 'N' {
					hour = 10 + int(s1[i]-'A')
					break
				} else if unicode.IsDigit(rune(s1[i])) {
					hour = int(s1[i] - '0')
					break
				}
			} else {
				if s1[i] >= 'A' && s1[i] <= 'G' {
					mask = true
					week = int(s1[i] - 'A')
				}
			}
		}
	}
	for i := 0; i < len(s3) && i < len(s4); i++ {
		if s3[i] == s4[i] && unicode.IsLetter(rune(s3[i])) {
			minute = i
			break
		}
	}
	fmt.Printf("%s %02d:%02d", WEEK[week], hour, minute)
}
