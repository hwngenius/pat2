package main

import "fmt"

func main() {
	var n, N int
	var max, sum int
	var i1, i2, j2 int
	var flag bool
	var first int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		if i == 0 {
			first = n
		}
		if flag == false {
			if n > 0 {
				sum = n
				flag = true
				i1 = n
				if sum > max {
					max = sum
					i2 = i1
					j2 = n
				}
			}
		} else {
			if sum += n; sum > max {
				max = sum
				i2 = i1
				j2 = n
			} else if sum < 0 {
				flag = false
			}
		}
	}
	if max != 0 {
		fmt.Println(max, i2, j2)
	} else {
		fmt.Println(0, first, n)
	}
}

// 25 22 wrong answer
