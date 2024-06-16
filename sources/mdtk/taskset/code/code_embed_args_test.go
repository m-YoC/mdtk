package code

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_escapeQuoteAndEnclose(t *testing.T) {
	tests := []struct {
		name string
		quote string
		str string
		expected string
	} {
		{"no quote", "", `Hello Go's Test`, `Hello Go's Test`},
		{"single quote, str has no quote", `'`, `Hello Test`, `'Hello Test'`},
		{"single quote, str has single quote", `'`, `Hello Go's Test`, `'Hello Go\'s Test'`},
		{"single quote, str has double quote", `'`, `Hello Go"s Test`, `'Hello Go"s Test'`},
		{"double quote, str has no quote", `"`, `Hello Test`, `"Hello Test"`},
		{"double quote, str has single quote", `"`, `Hello Go's Test`, `"Hello Go's Test"`},
		{"double quote, str has double quote", `"`, `Hello Go"s Test`, `"Hello Go\"s Test"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := escapeQuoteAndEnclose(tt.str, tt.quote, `\`)
			assert.Equal(t, tt.expected, s)
		})
	}

}


