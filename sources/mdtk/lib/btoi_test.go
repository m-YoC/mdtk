package lib

import (
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
)

func testBtoiBase[T Integer](t *testing.T) {
	t.Run("Btoi[" + reflect.TypeOf(T(0)).String() + "]", func(t *testing.T) {
		// assert.Equal(t, reflect.TypeOf(T(0)).String(), reflect.TypeOf(Btoi[T](true)).String())
		assert.IsType(t,T(0), Btoi[T](true))
		assert.Equal(t, T(1), Btoi[T](true))
		assert.Equal(t, T(0), Btoi[T](false))
	})
} 

func Test_Btoi(t *testing.T) {

	testBtoiBase[int](t)
	testBtoiBase[uint](t)
	testBtoiBase[int64](t)
	testBtoiBase[uint64](t)
	
}
