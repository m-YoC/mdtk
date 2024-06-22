package lib

import (
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ArraySlice(t *testing.T) {
	array := []string{"hoge", "huga", "piyo", "yahooooo"}

	t.Run("Have", func(t *testing.T){
		assert.True(t, Slice(array).Have("hoge"))
		assert.True(t, Slice(array).Have("huga"))
		assert.True(t, Slice(array).Have("piyo"))

		assert.False(t, Slice(array).Have("yahoo"))
		assert.False(t, Slice(array).Have("fizz"))
		assert.False(t, Slice(array).Have("bizz"))
	})

	t.Run("HaveFunc", func(t *testing.T){
		d, b := Slice(array).HaveFunc(func(d string)bool { return strings.HasPrefix(d, "hu") })
		assert.True(t, b)
		assert.Equal(t, "huga", d)
		_, b = Slice(array).HaveFunc(func(d string)bool { return strings.HasPrefix(d, "he") })
		assert.False(t, b)
	})
}


func Test_VarSlice(t *testing.T) {
	array := []string{"hoge", "huga", "piyo", "yahooooo"}

	t.Run("IsContainedIn", func(t *testing.T){
		assert.True(t, Var("hoge").IsContainedIn(array))
		assert.True(t, Var("huga").IsContainedIn(array))
		assert.True(t, Var("piyo").IsContainedIn(array))
		
		assert.False(t, Var("yahoo").IsContainedIn(array))
		assert.False(t, Var("fizz").IsContainedIn(array))
		assert.False(t, Var("bizz").IsContainedIn(array))
	})

	
}


