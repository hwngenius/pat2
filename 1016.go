package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type myslice []string

func (m myslice) Len() int {
	return len(m)
}

func (m myslice) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

var price [25]float64
var day float64

func cost(s1, s2 string) (item float64, minute int) {
	d1, _ := strconv.Atoi(s1[:2])
	d2, _ := strconv.Atoi(s2[0:2])
	h1, _ := strconv.Atoi(s1[3:5])
	h2, _ := strconv.Atoi(s2[3:5])
	m1, _ := strconv.Atoi(s1[6:])
	m2, _ := strconv.Atoi(s2[6:])
	for d1 < d2 || h1 < h2 || m1 < m2 {
		minute++
		item += price[h1]
		m1++
		if m1 >= 60 {
			m1 = 0
			h1++
		}
		if h1 >= 24 {
			h1 = 0
			d1++
		}
	}
	return item, minute
}
func main() {
	online := make(map[string][]string)
	offline := make(map[string][]string)
	namemask := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	L := strings.Split(s, " ")
	for i := 0; i < len(L); i++ {
		t, _ := strconv.Atoi(L[i])
		price[i] = float64(t)
		day += price[i]
	}
	input.Scan()
	s = input.Text()
	for input.Scan() {
		s = input.Text()
		L = strings.Split(s, " ")
		namemask[L[0]] = true
		if L[2] == "on-line" {
			online[L[0]] = append(online[L[0]], L[1])
		} else {
			offline[L[0]] = append(offline[L[0]], L[1])
		}
	}
	var name []string
	for k, _ := range namemask {
		name = append(name, k)
	}
	for _, v := range online {
		sort.Sort(myslice(v))
	}
	for _, v := range offline {
		sort.Sort(myslice(v))
	}
	sort.Sort(myslice(name))
	for i := 0; i < len(name); i++ {
		var j, k = 0, 0
		var total float64
		var mask int
		for ; j < len(offline[name[i]]); j++ {
			for j < len(offline[name[i]]) && k < len(online[name[i]]) && offline[name[i]][j] > online[name[i]][k] {
				mask++
				if mask == 1 {
					fmt.Printf("%s %s\n", name[i], offline[name[i]][j][:2])
				}
				k++
			}
			item, minute := cost(online[name[i]][k-1][3:], offline[name[i]][j][3:])
			fmt.Printf("%s %s %d $%.2f\n", online[name[i]][k-1][3:], offline[name[i]][j][3:], minute, item/100.0)
			total += item
		}
		fmt.Printf("Total amount: $%.2f\n", total/100.0)
	}
}

// 25 15 no-zero return value
