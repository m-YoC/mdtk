package parse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_LexArgString(t *testing.T) {

	t.Run("positive", func(t *testing.T) {
		actual, err := LexArgString("echo this\\\" is 'test string'")
		expected := []string{"echo", "this\\\"", "is", "'test string'"}
	
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("negative: escape sequence error", func(t *testing.T) {
		tests := []struct {
			name string
			args string
		}{
			{"space after \\", "echo this is\\ 'test string'"},
			{"EOF after \\", "echo this is 'test string'\\"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := LexArgString(tt.args)
				assert.Error(t, err)
			})
		}

	})

	t.Run("negative: quote error", func(t *testing.T) {
		tests := []struct {
			name string
			args string
		}{
			{"string quotes not closed", "echo this is 'test string"},
			{"string double-quotes not closed", `echo this is "test string`},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := LexArgString(tt.args)
				assert.Error(t, err)
			})
		}

	})
}