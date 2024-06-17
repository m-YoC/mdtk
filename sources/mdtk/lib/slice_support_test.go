package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_VarSlice(t *testing.T) {
	array := []string{"hoge", "huga", "piyo", "yahooooo"}

	assert.True(t, Var("hoge").IsContainedIn(array))
	assert.True(t, Var("huga").IsContainedIn(array))
	assert.True(t, Var("piyo").IsContainedIn(array))
	assert.False(t, Var("yahoo").IsContainedIn(array))
	assert.False(t, Var("fizz").IsContainedIn(array))
	assert.False(t, Var("bizz").IsContainedIn(array))
}


