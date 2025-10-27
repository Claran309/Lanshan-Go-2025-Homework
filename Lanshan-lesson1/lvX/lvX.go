package main

import (
	"fmt"
)

func average(sum, n int) float64 {
	return float64(sum) / float64(n)
}

var n, sum int

func main() {
	for num := 1; num != 0; n, sum = n+1, sum+num {
		fmt.Print("请输入一个整数(输入0结束): ")
		fmt.Scanln(&num)
	}
	if ave := average(sum, n-1); ave >= 60 {
		fmt.Printf("平均成绩为 %.2f，成绩合格", ave)
	} else {
		fmt.Printf("平均成绩为 %.2f，成绩不合格", ave)
	}
}
