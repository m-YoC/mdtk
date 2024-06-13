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

func Test_Example(t *testing.T) {
	t.Run("desc", func(t *testing.T) {
		tests := TestCases[string, string] {
			{Name: "test", TestArg: "hello", Expected: "hello"},
		}
		
		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				assert.Equal(t, tt.Expected, tt.TestArg)
			})
		}
	})
}
