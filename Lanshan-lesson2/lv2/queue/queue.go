package queue

type Queue struct {
	elements []interface{} //0永远为Top/Front
}

func (q *Queue) Push(item interface{}) {
	q.elements = append(q.elements, item)
}

func (q *Queue) Pop() {
	if !q.Empty() {
		q.elements = q.elements[1:] //delete top
	}
}

func (q *Queue) Front() interface{} {
	if q.Empty() {
		return nil //队列空返回nil
	}
	item := q.elements[0]
	return item
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Clear() {
	q.elements = make([]interface{}, 0)
}
