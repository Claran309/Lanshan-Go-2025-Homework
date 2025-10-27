package main

import (
	"fmt"
)

var n, sum int

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println(sum)
}
