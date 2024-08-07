package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	Positive = "positive"
	Negative = "negative"
)

type TestCase[TA, TE any] struct {
	Name string
	TestArg TA
	Expected TE
}

type TestCases[TA, TE any] []TestCase[TA, TE]

func (tests TestCases[TA, TE]) Get(i int) TestCase[TA, TE] {
	return tests[i]
}

func (tests TestCases[TA, TE]) Run(t *testing.T, f func(*testing.T, int)) {
	t.Helper()

	for i, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Helper()
			f(t, i)
		})
	}
}

func Test_Example(t *testing.T) {
	t.Run("desc", func(t *testing.T) {
		tests := TestCases[string, string] {
			{Name: "test", TestArg: "hello", Expected: "hello"},
		}

		tests.Run(t, func(t *testing.T, i int) {
			tt := tests.Get(i)
			assert.Equal(t, tt.Expected, tt.TestArg)
		})
	})
}
