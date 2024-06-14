package base

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_NewExit(t *testing.T) {
	test_status := 0
	old := osExit
	defer func() { osExit = old }()
	osExit = func(code int) { test_status = 10 }

	// First is zero.
	assert.Equal(t, 0, test_status)

	internal := func() {
		// External osExit function always sets 10.
		osExit(20)
		assert.Equal(t, 10, test_status)

		// Set new internal osExit and reset after defer.
		defer NewExit(func(status int) { test_status = status })()
		
		assert.Equal(t, 10, test_status)
		osExit(20)
		assert.Equal(t, 20, test_status)
	}

	internal()

	// after internal-defer
	assert.Equal(t, 20, test_status)
	// Did osExit function revert to the external one?
	osExit(20)
	assert.Equal(t, 10, test_status)

} 

func Test_AddFinalize(t *testing.T) {
	assert.Equal(t, 0, len(finalize))
	AddFinalize(func(){ /* Nothing */ })
	assert.Equal(t, 1, len(finalize))
	finalize = []func(){}
	assert.Equal(t, 0, len(finalize))
}

func Test_MdtkExit(t *testing.T) {
	test_status := -1
	defer NewExit(func(status int) { test_status = status })()
	assert.Equal(t, -1, test_status)

	called_num := 0
	AddFinalize(func(){ called_num++ })
	AddFinalize(func(){ called_num++ })
	AddFinalize(func(){ called_num++ })
	assert.Equal(t, 0, called_num)

	MdtkExit(10)

	assert.Equal(t, 10, test_status)
	assert.Equal(t, 3, called_num)
}


