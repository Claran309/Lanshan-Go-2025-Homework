package stack

type Stack struct {
	//切片储存栈元素，i=len(elements)为top
	elements []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() {
	if !s.Empty() {
		lastIdx := len(s.elements) - 1
		s.elements = s.elements[:lastIdx]
	}
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Clear() {
	s.elements = make([]interface{}, 0)
}
