package args

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Arg(t *testing.T) {
	const (
		positive = "positive"
		negative = "negative"
	)
	
	t.Run("Validate", func(t *testing.T) {
		tests := []struct {
			aname string
			expected string
		} {
			{"key=value", positive},
			{"_key=value", positive},
			{"key=va lu e", positive},
			{"key='va lu e'", positive},
			{"key =value", negative},
			{"ke~y=value", negative},
			{"key: value", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %v", tt.aname, tt.expected), func(t *testing.T) {
				err := Arg(tt.aname).Validate()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
			})
		}
	})

	t.Run("GetData", func(t *testing.T) {
		tests := []struct {
			aname string
			expected_k string
			expected_v string
			expected_err string
		} {
			{"key=value", "key", "value", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %v and get key:%s, value:%s when positive", 
					tt.aname, tt.expected_err, tt.expected_k, tt.expected_v), func(t *testing.T) {
				k, v, err := Arg(tt.aname).GetData()
				if tt.expected_err == positive {
					if assert.NoError(t, err) {
						assert.Equal(t, tt.expected_k, k)
						assert.Equal(t, tt.expected_v, v)
					}
				} else {
					assert.Error(t, err)
				}
			})
		}

	})

}


func Test_Args(t *testing.T) {
	const (
		positive = "positive"
		negative = "negative"
	)
	
	t.Run("Validate", func(t *testing.T) {
		tests := []struct {
			anames []string
			expected string
		} {
			{[]string{"key=value", "key=value", "key=value"}, positive},
			{[]string{"key=value", "key: value", "key=value"}, positive},
			{[]string{"key=value", "key; value", "key=value"}, negative},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("AND of '%v' is %v", tt.anames, tt.expected), func(t *testing.T) {
				err := ToArgs(tt.anames...).Validate()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
			})
		}
	})

}

