package main

import (
	"fmt"
	"sort"
)

var now int

type student struct {
	id    string
	name  string
	grade int
}
type myslice []student

func (m myslice) Len() int {
	return len(m)
}
func (m myslice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m myslice) Less(i, j int) bool {
	if now == 1 {
		return m[i].id < m[j].id
	} else {
		if now == 2 {
			if m[i].name != m[j].name {
				return m[i].name < m[j].name
			} else {
				return m[i].id < m[j].id
			}
		} else {
			if m[i].grade != m[j].grade {
				return m[i].grade < m[j].grade
			} else {
				return m[i].id < m[j].id
			}
		}
	}
}
func main() {
	var N, C int
	var sample student
	var input []student
	fmt.Scanf("%d %d", &N, &C)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %s %d", &sample.id, &sample.name, &sample.grade)
		input = append(input, sample)
	}
	now = C
	sort.Sort(myslice(input))
	for i := 0; i < len(input); i++ {
		fmt.Printf("%s %s %d\n", input[i].id, input[i].name, input[i].grade)
	}
}

// 25/21 out of time
