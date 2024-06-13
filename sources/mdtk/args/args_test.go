package args

import (
	"mdtk/lib"
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	positive = "positive"
	negative = "negative"
)

func Test_Arg(t *testing.T) {
	t.Run("Validate", func(t *testing.T) {
		tests := lib.TestCases[string, string] {
			{Name: "Alphabet=Alphabet is ok", TestArg: "key=value", Expected: positive},
			{Name: "Underbar is also ok", TestArg: "_key=value", Expected: positive},
			{Name: "In value, all chars are ok", TestArg: "key=va &lu! e", Expected: positive},
			{Name: "Can enclose value in quotes", TestArg: "key='va lu e'", Expected: positive},
			{Name: "Instead of '=', ':' is also ok", TestArg: "key:value", Expected: positive},
			{Name: "Between key and '=', have not to set space", TestArg: "key =value", Expected: negative},
			{Name: "Key's chars are only alphabet and underbar", TestArg: "ke~y=value", Expected: negative},
			{Name: "Key's chars are only alphabet and underbar", TestArg: "ke y=value", Expected: negative},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				err := Arg(tt.TestArg).Validate()
				if tt.Expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
			})
		}
	})

	t.Run("GetData", func(t *testing.T) {
		type ExpectedT struct {
			Key string
			Value string
			PN string
		}
		tests := lib.TestCases[string, ExpectedT] {
			{Name: "Basic", TestArg: "key=value", Expected: ExpectedT{Key: "key", Value: "value", PN: positive}},
			{Name: "Key's chars are only alphabet and underbar", TestArg: "ke y=value", Expected: ExpectedT{PN: negative}},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				k, v, err := Arg(tt.TestArg).GetData()
				if tt.Expected.PN == positive {
					if assert.NoError(t, err) {
						assert.Equal(t, tt.Expected.Key, k)
						assert.Equal(t, tt.Expected.Value, v)
					}
				} else {
					assert.Error(t, err)
				}
			})
		}

	})

}


func Test_Args(t *testing.T) {	
	t.Run("Validate", func(t *testing.T) {
		tests := lib.TestCases[[]string, string] {
			{Name: "All is ok", TestArg: []string{"key=value", "key=value", "key=value"}, Expected: positive},
			{Name: "All is ok", TestArg: []string{"key=value", "key: value", "key=value"}, Expected: positive},
			{Name: "Includes ng", TestArg: []string{"key=value", "key; value", "key=value"}, Expected: negative},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				err := ToArgs(tt.TestArg...).Validate()
				if tt.Expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
			})
		}
	})

}

