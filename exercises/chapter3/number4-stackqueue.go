package chapter3

type Queue interface {
	push(int)
	pop() int
	size() int
}

func NewStackQueue() Queue {
	return &StackQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

type StackQueue struct {
	inStack  []int
	outStack []int
}

func (q *StackQueue) push(n int) {
	q.inStack = append(q.inStack, n)
}

func (q *StackQueue) pop() int {
	if len(q.outStack) > 0 {
		outVal := q.outStack[len(q.outStack)-1]
		q.outStack = q.outStack[:len(q.outStack)-1]
		return outVal
	} else {

	}
	return 0
}

func (q *StackQueue) size() int {
	return 0
}
