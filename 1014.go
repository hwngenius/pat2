package main

import "fmt"

type queue struct {
	element     [10000]int
	front, back int
}

func (q *queue) len() int {
	return q.back - q.front
}

func (q *queue) pushback(e int) {
	q.element[q.back] = e
	q.back++
}

func (q *queue) popfront() int {
	q.front++
	return q.element[q.front-1]
}

func (q *queue) passtime(e int) {
	q.element[q.front] -= e
}

func (q *queue) fisrttime() int {
	return q.element[q.front]
}

func main() {
	var windwosqueue [20]queue
	var windwostime [20]int
	var N, M, K, Q int
	var finish [10000]int
	var time, count int
	fmt.Scanf("%d %d %d %d", &N, &M, &K, &Q)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Scan(&time)
			windwosqueue[j].pushback(time)
			if windwostime[j] < (17-8)*60 {
				windwostime[j] += time
				finish[count] = windwostime[j]
			} else {
				windwostime[j] += time
				finish[count] = -1
			}
			count++
		}
	}

	for i := N * M; i < K; i++ {
		var min = 1000000000000
		var index = 0
		fmt.Scan(&time)
		for j := 0; j < N; j++ {
			// fmt.Println(windwosqueue[j].fisrttime())
			if t := windwosqueue[j].fisrttime(); t < min {
				index = j
				min = t
				// fmt.Println(j, t)
			}
		}
		for j := 0; j < N; j++ {
			windwosqueue[j].passtime(min)
		}
		// fmt.Println()
		windwosqueue[index].popfront()
		windwosqueue[index].pushback(time)
		if windwostime[index] < (17-8)*60 {
			windwostime[index] += time
			finish[count] = windwostime[index]
		} else {
			finish[count] = -1
		}
		count++
	}
	for i := 0; i < Q; i++ {
		fmt.Scan(&time)
		if finish[time-1] == -1 {
			fmt.Println("Sorry")
			continue
		}

		fmt.Printf("%02d:%02d\n", (finish[time-1]/60+8)%24, finish[time-1]%60)
	}
}

// 25 23 wrong answer
