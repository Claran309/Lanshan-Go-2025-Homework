package lv2_func

import (
	"Lesson_1/Lanshan-lesson2/lv2/queue"
	"Lesson_1/Lanshan-lesson2/lv2/stack"
	"strconv"
	"unicode"
)

var priority = map[rune]int{
	'*': 2,
	'/': 2,
	'+': 1,
	'-': 1,
}

func InToPost(in string, s *stack.Stack, q *queue.Queue) {
	s.Clear()
	q.Clear()
	for i := 0; i < len(in); i++ {
		ch := rune(in[i])
		switch {
		case unicode.IsDigit(ch):
			endIdx := i
			for j := i; j < len(in) && (in[j] == '.' || (in[j] >= '0' && in[j] <= '9')); j++ {
				endIdx = j
			}
			num := in[i : endIdx+1]
			i = endIdx
			n, err := strconv.ParseFloat(num, 64)
			if err != nil {
			}
			q.Push(n)
		case ch == '(':
			s.Push(ch)
		case ch == ')':
			check := s.Top().(rune)
			for check != '(' && !s.Empty() {
				q.Push(check)
				s.Pop()
				check = s.Top().(rune)
			}
			s.Pop()
		case ch == '*' || ch == '/' || ch == '+' || ch == '-':
			for !s.Empty() && priority[s.Top().(rune)] >= priority[ch] {
				ru := s.Top().(rune)
				q.Push(ru)
				s.Pop()
			}
			s.Push(ch)
		}
	}
	for !s.Empty() {
		q.Push(s.Top())
		s.Pop()
	}
}

func PostCount(q *queue.Queue) float64 {
	var s stack.Stack
	var n1, n2, ans float64
	for !q.Empty() {
		topElement := q.Front()
		q.Pop()
		switch v := topElement.(type) {
		case float64:
			s.Push(v)
		case rune:
			n2 = s.Top().(float64)
			s.Pop()
			n1 = s.Top().(float64)
			s.Pop()
			switch v {
			case '+':
				ans = n1 + n2
			case '-':
				ans = n1 - n2
			case '*':
				ans = n1 * n2
			case '/':
				ans = n1 / n2
			}
			s.Push(ans)
		}
	}
	result := s.Top().(float64)
	s.Clear()
	return result
}

func Run(input string) float64 {
	var q queue.Queue
	var s stack.Stack
	InToPost(input, &s, &q)
	return PostCount(&q)
}
