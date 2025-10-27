package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	if n == 1 {
		return big.NewInt(1)
	}
	res := big.NewInt(int64(n))
	return res.Mul(res, factorial(n-1))
}

var num int

func main() {
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}
