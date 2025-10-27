package main

import (
	"fmt"
)

const pi float32 = 3.14

func main() {
	var r int
	fmt.Scan(&r)
	fmt.Printf("area = %f", float32(r)*float32(r)*pi)
}
