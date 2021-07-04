package chapter3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceQueue(t *testing.T) {
	q := NewSliceQueue()
	testcase := []int{2, 4, 6, 8, 10}
	for _, n := range testcase {
		q.push(n)
	}
	for _, want := range testcase {
		got := q.pop()
		assert.Equal(t, want, got)
	}
	assert.Equal(t, 0, q.size())

}
