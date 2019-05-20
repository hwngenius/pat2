package main

import "fmt"

type item struct {
	xishu  float64
	zhishu int
}

func main() {
	var a, b, c [20]item
	var an, bn int
	var xishu float64
	var zhishu int
	fmt.Scan(&an)
	for i := 0; i < an; i++ {
		fmt.Scanf("%d %f", &zhishu, &xishu)
		a[i].zhishu = zhishu
		a[i].xishu = xishu
	}
	fmt.Scan(&bn)
	for i := 0; i < bn; i++ {
		fmt.Scanf("%d %f", &zhishu, &xishu)
		b[i].zhishu = zhishu
		b[i].xishu = xishu
	}
	var i, j, k int
	for i, j, k = 0, 0, 0; j < an && k < bn; i++ {
		if a[j].zhishu == b[k].zhishu {
			if a[j].xishu == -b[k].xishu {
				i--
				j++
				k++
				continue
			}
			c[i].zhishu = a[j].zhishu
			c[i].xishu = a[j].xishu + b[k].xishu
			j++
			k++
		} else if a[j].zhishu > b[k].zhishu {
			c[i] = a[j]
			j++
		} else {
			c[i] = b[k]
			k++
		}
	}
	for ; j < an; j++ {
		c[i] = a[j]
		i++
	}
	for ; k < bn; k++ {
		c[i] = b[k]
		i++
	}
	fmt.Printf("%d", i)
	for i1 := 0; i1 < i; i1++ {
		fmt.Printf(" %d %.1f", c[i1].zhishu, c[i1].xishu)
	}
}
