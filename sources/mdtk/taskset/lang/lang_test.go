package lang

import (
	"mdtk/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)




func Test_removeOpC(t *testing.T) {
	type S = []string
	tests := lib.TestCases[S, S] {
		{Name: "Only has -c", TestArg: S{"-c"}, Expected: S{}},
		{Name: "Has -cx", TestArg: S{"-cx"}, Expected: S{"-x"}},
		{Name: "Has -a -cx", TestArg: S{"-a", "-cx"}, Expected: S{"-a", "-x"}},
		{Name: "Has -cx -a", TestArg: S{"-cx", "-a"}, Expected: S{"-cx", "-a"}},
		{Name: "Only has -x", TestArg: S{"-x"}, Expected: S{"-x"}},
		{Name: "Has -Command", TestArg: S{"-Command"}, Expected: S{}},
		{Name: "Has -Command", TestArg: S{"-a", "-Command"}, Expected: S{"-a"}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			actual := removeOpC(tt.TestArg)
			assert.Equal(t, tt.Expected, actual)
		})
	}
}

