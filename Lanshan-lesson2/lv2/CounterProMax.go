// 我直接搬的我在隔壁写的计算器作业了
// 代码，包和注释（还有调试过程QWQ）都是手敲的（不然我都不知道自己写到哪了），不信任我可以拷问我QWQ
package main

import (
	"Lesson_1/Lanshan-lesson2/lv2/queue"
	"Lesson_1/Lanshan-lesson2/lv2/stack"
	"fmt"
	"strconv"
	"unicode"
)

var priority = map[rune]int{
	'*': 2,
	'/': 2,
	'+': 1,
	'-': 1,
}

func InToPost(in string, s *stack.Stack, q *queue.Queue) { // 中缀表达式转后缀表达式
	s.Clear()
	q.Clear()
	for i := 0; i < len(in); i++ {
		ch := rune(in[i])
		switch {
		case unicode.IsDigit(ch): // 如果是数字
			endIdx := i
			for j := i; j < len(in) && (in[j] == '.' || (in[j] >= '0' && in[j] <= '9')); j++ { // 扫描数字范围
				endIdx = j // 标记endIndex
			}
			num := in[i : endIdx+1]               // 储存数字
			i = endIdx                            // 更新下标
			n, err := strconv.ParseFloat(num, 64) // 字符串转浮点数
			if err != nil {
			}
			q.Push(n) //将数字插入队列
		case ch == '(': // 如果是左括号，则将左括号压入栈
			s.Push(ch)
		case ch == ')': // 如果是右括号，则弹出左括号前的所有栈元素
			check := s.Top().(rune)
			for check != '(' && !s.Empty() { // 弹出左括号前的所有栈元素
				//fmt.Printf("---------ru = %c------------\n", ru)
				q.Push(check) // 将运算符插入后缀表达式
				s.Pop()
				check = s.Top().(rune)
			}
			s.Pop() // 弹出左括号
		case ch == '*' || ch == '/' || ch == '+' || ch == '-':
			for !s.Empty() && priority[s.Top().(rune)] >= priority[ch] { // 当栈不为空且栈顶运算符优先级 >= 当前运算符优先级
				ru := s.Top().(rune)
				q.Push(ru)
				s.Pop()
			}
			s.Push(ch) // 将当前运算符压入栈
		}
	}
	for !s.Empty() { // 向队列插入剩余运算符
		q.Push(s.Top())
		s.Pop()
	}
}

func PostCount(q *queue.Queue) float64 {
	var s stack.Stack
	var n1, n2, ans float64
	for !q.Empty() { // 扫描队列元素
		topElement := q.Front()
		q.Pop()
		switch v := topElement.(type) {
		case float64: // 如果是数字
			s.Push(v)
		case rune: // 如果是操作符
			n2 = s.Top().(float64)
			s.Pop()
			n1 = s.Top().(float64)
			s.Pop()    // 弹出两个元素
			switch v { // 运算
			case '+':
				ans = n1 + n2
			case '-':
				ans = n1 - n2
			case '*':
				ans = n1 * n2
			case '/':
				ans = n1 / n2
			}
			s.Push(ans) // 将结果压回栈内
		}
	}
	result := s.Top().(float64) // 栈顶元素为结果
	s.Clear()
	return result
}

func main() {
	fmt.Printf("桓因使用Go语言计算器！\n请输入一个合法的算数表达式，例如:(3.14+2.71)*2/5\n输入exit退出程序")
	var input string
	for {
		fmt.Printf("\n请输入:")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
		var q queue.Queue
		var s stack.Stack
		InToPost(input, &s, &q) // 中缀表达式转后缀表达式
		/*for !q.Empty() {                          debug遗留输出
			value, ok := q.Front().(rune)
			value2, ok2 := q.Front().(float64)
			if ok {
				fmt.Printf("%c ", value)
			} else if ok2 {
				fmt.Printf("%f ", value2)
			}
			q.Pop()
		}*/
		fmt.Printf("结果是%f\n", PostCount(&q)) // 后缀表达式求值
	}
	fmt.Printf("感谢使用！再见！\n")
}
