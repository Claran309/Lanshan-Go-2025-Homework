package main

import (
	"fmt"
)

var n int

func main() {
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < n; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Print(a[i], " ")
	}
}
