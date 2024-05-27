package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	var stack Stack[int]

	before := func() {
		stack = Stack[int]{}
	}

	t.Run("check", func(t *testing.T) {
		before()
		stack.Push(1)

		assert.Equal(t, 1, stack.Size())
		{
			a, ok := stack.Top()
			assert.Equal(t, 1, a)
			assert.True(t, ok)
		}
		
		{
			poped, ok := stack.Pop()
			assert.Equal(t, 1, poped)
			assert.True(t, ok)
		}

		assert.Equal(t, 0, stack.Size())
		
		{
			a, ok := stack.Top()
			assert.Zero(t, a)
			assert.False(t, ok)
		}

		{
			poped, ok := stack.Pop()
			assert.Zero(t, poped)
			assert.False(t, ok)
		}
	})

}