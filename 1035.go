package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func myFun(b byte) (byte, bool) {
	switch b {
	case '0':
		return '%', true
	case 'l':
		return 'L', true
	case 'O':
		return 'o', true
	case '1':
		return '@', true
	default:
		return b, false
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	var ans []string
	var sum int
	N, _ := strconv.Atoi(input.Text())
	for input.Scan() {
		L := strings.Split(input.Text(), " ")
		var b []byte
		mask := false
		for i := 0; i < len(L[1]); i++ {
			t1, t2 := myFun(L[1][i])
			if t2 == true {
				mask = true
			}
			b = append(b, t1)
		}
		if mask == true {
			sum++
			ans = append(ans, L[0]+" "+string(b))
		}
	}
	if ans == nil {
		if N == 1 {
			fmt.Println("There is 1 account and no account is modified")
		} else {
			fmt.Printf("There are %d accounts and no account is modified\n", N)
		}
	} else {
		fmt.Println(sum)
		for i := 0; i < len(ans); i++ {
			fmt.Println(ans[i])
		}
	}
}
