package main

import "fmt"

func main() {
	var set = [...]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C'}
	fmt.Printf("#")
	var n int
	for i := 0; i < 3; i++ {
		fmt.Scan(&n)
		fmt.Printf("%c", set[n/13])
		fmt.Printf("%c", set[n%13])
	}
}
