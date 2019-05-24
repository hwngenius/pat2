package main

import (
	"fmt"
	"sort"
)

type arrive struct {
	hh      int
	mm      int
	ss      int
	process int
}

type myslice []arrive

func (m myslice) Len() int {
	return len(m)
}

func (m myslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myslice) Less(i, j int) bool {
	if m[i].hh != m[j].hh {
		return m[i].hh < m[j].hh
	} else if m[i].mm != m[j].mm {
		return m[i].mm < m[j].mm
	} else {
		return m[i].ss < m[j].ss
	}
}

func change(person arrive) int {
	return person.hh*60*60 + person.mm*60 + person.ss
}

func unchange(n int) (h, m, s int) {
	return n / 3600, n % 3600 / 60, n % 60
}
func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)
	begin := change(arrive{8, 0, 0, 0})
	end := change(arrive{17, 0, 0, 0})
	var windwosTime [100]int
	var input []arrive
	var sample arrive
	var wait int
	var n int
	for i := 0; i < N; i++ {
		fmt.Scanf("%d:%d:%d %d", &sample.hh, &sample.mm, &sample.ss, &sample.process)
		if sample.process > 60 {
			sample.process = 60
		}
		input = append(input, sample)
	}
	sort.Sort(myslice(input))
	// for i := 0; i < len(input); i++ {
	// 	fmt.Println(input[i])
	// }
	for i := 0; i < K && i < N; i++ {
		if t1, t2 := change(input[i]), begin; t1 <= t2 {
			wait += t2 - t1
			n++
			windwosTime[i] = begin + input[i].process*60
		} else if t1 >= end {
			continue
		} else {
			windwosTime[i] = t1 + input[i].process*60
		}
		// fmt.Println(unchange(windwosTime[i]))
	}
	j := 0
	for i := K; i < N; i++ {
		min := 10000000000
		for k := 0; k < K; k++ {
			if windwosTime[k] < min {
				min = windwosTime[k]
				j = k
			}
		}
		if t1 := change(input[i]); t1 >= end {
			break
		} else if t1 <= windwosTime[j] {
			wait += windwosTime[j] - t1
			n++
			windwosTime[j] += input[i].process * 60
		} else {
			windwosTime[j] = t1 + input[i].process*60
		}
		// fmt.Printf("%d ", j)
		// fmt.Println(unchange(windwosTime[j]))
	}
	fmt.Println(wait, n)
	if wait != 0 {
		fmt.Printf("%.1f", float64(wait)/float64(n)/float64(60))
	} else {
		fmt.Println("0.0")
	}
}
