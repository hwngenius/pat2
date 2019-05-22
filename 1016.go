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

var price [24]float64
var day float64

func cost(s1, s2 string) (item float64) {
	d1, _ := strconv.Atoi(s1[:2])
	d2, _ := strconv.Atoi(s2[0:2])
	item += float64(d2-d1) * day
	h1, _ := strconv.Atoi(s1[3:5])
	h2, _ := strconv.Atoi(s2[3:5])
	m1, _ := strconv.Atoi(s1[6:])
	m2, _ := strconv.Atoi(s2[6:])
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
			t1, _ := strconv.Atoi(online[name[i]][k-1][3:5])
			t2, _ := strconv.Atoi(offline[name[i]][j][3:5])
			d := t2 - t1
			t1, _ = strconv.Atoi(online[name[i]][k-1][6:8])
			t2, _ = strconv.Atoi(offline[name[i]][j][6:8])
			fmt.Println(t1, t2)
			t := 0
			h := t2 - t1
			t1, _ = strconv.Atoi(online[name[i]][k-1][9:])
			t2, _ = strconv.Atoi(offline[name[i]][j][9:])
			m := t2 - t1
			item := cost(offline[name[i]][k-1][3:], offline[name[i]][j][3:])
			minute := d*24*60 + h*60 + m
			fmt.Printf("%s %s %d $%.2f\n", online[name[i]][k-1][3:], offline[name[i]][j][3:], minute, item)
			total += item
		}
		fmt.Printf("Total amount: $%.2f\n", total)
	}
}
