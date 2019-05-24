package main

import (
	"fmt"
)

type Tree struct {
	elment      int
	left, right *Tree
}

func myfunc(last, mid []int, T *Tree) *Tree {
	l := 0
	if l = len(last); l != len(mid) {
		fmt.Println("SB")
		return nil
	} else if l == 0 {
		return nil
	}
	if T == nil {
		i := 0
		T = new(Tree)
		T.elment = last[l-1]
		for i = 0; i < l; i++ {
			if mid[i] == last[l-1] {
				break
			}
		}
		T = new(Tree)
		T.elment = last[l-1]
		T.left = myfunc(last[:i], mid[:i], T.left)
		T.right = myfunc(last[i:l-1], mid[i+1:], T.right)
	}
	return T
}

type queue struct {
	element     [32]*Tree
	front, back int
}

func (q *queue) pushback(T *Tree) {
	q.element[q.back] = T
	q.back++
}
func (q *queue) popfront() (T *Tree) {
	q.front++
	return q.element[q.front-1]
}
func (q *queue) len() int {
	return q.back - q.front
}

func main() {
	var T *Tree
	var N, n int
	var last []int
	var mid []int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		last = append(last, n)
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		mid = append(mid, n)
	}
	var result []int
	T = myfunc(last, mid, T)
	var q queue
	q.pushback(T)
	for q.len() != 0 {
		e := q.popfront()
		result = append(result, e.elment)
		if e.left != nil {
			q.pushback(e.left)
		}
		if e.right != nil {
			q.pushback(e.right)
		}
	}
	for i := 0; i < len(result); i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", result[i])
	}
}
