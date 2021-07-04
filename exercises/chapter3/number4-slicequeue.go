package chapter3

type SliceQueue struct {
	slice []int
}

func NewSliceQueue() Queue {
	return &SliceQueue{slice: []int{}}
}

func (q *SliceQueue) push(n int) {
	q.slice = append(q.slice, n)
}

func (q *SliceQueue) pop() int {
	popVal := q.slice[0]
	q.slice = q.slice[1:]
	return popVal
}

func (q *SliceQueue) size() int {
	return len(q.slice)
}
