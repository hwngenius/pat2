package main

import (
	"fmt"
	"sort"
)

var now int

type Student struct {
	id    int
	grade [4]int
}

type myslice []Student

func (m myslice) Len() int {
	return len(m)
}
func (m myslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m myslice) Less(i, j int) bool {
	return m[i].grade[now] > m[j].grade[now]
}

func main() {
	var change = [...]byte{'A', 'C', 'M', 'E'}
	var N, M int
	var m [1000000][4]int
	var sample Student
	var input []Student
	fmt.Scan(&N)
	fmt.Scan(&M)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d %d %d", &sample.id, &sample.grade[1], &sample.grade[2], &sample.grade[3])
		sample.grade[0] = sample.grade[1] + sample.grade[2] + sample.grade[3]
		input = append(input, sample)
	}
	for now = 0; now < 4; now++ {
		sort.Sort(myslice(input))
		t := 1
		m[input[0].id][now] = 1
		for j := 1; j < N; j++ {
			if input[j].grade[now] == input[j-1].grade[now] {
				m[input[j].id][now] = m[input[j-1].id][now]
				t++
			} else {
				m[input[j].id][now] = m[input[j-1].id][now] + t
				t = 1
			}
		}
	}
	var query int

	for i := 0; i < M; i++ {
		fmt.Scan(&query)
		min := N + 1
		MinIndex := 0
		for index, rank := range m[query] {
			if rank < min {
				min = rank
				MinIndex = index
			}
		}
		if min == 0 {
			fmt.Println("N/A")
			continue
		}
		fmt.Printf("%d %c\n", min, change[MinIndex])
	}
}
