package chapter3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiset(t *testing.T) {
	m := NewMultiSet()
	assert.Equal(t, "{ }", String(m))
	assert.Equal(t, 0, Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	assert.Equal(t, "{3 3 3 3}", String(m))
	assert.Equal(t, 4, Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	assert.Equal(t, 3, Count(m, "3"))
	assert.Equal(t, 1, Count(m, "1"))
	assert.Equal(t, 1, Count(m, "2"))
	assert.Equal(t, 0, Count(m, "5"))

}
