package group

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Group(t *testing.T) {
	
	t.Run("IsPrivate", func(t *testing.T) {
		tests := []struct {
			gname string
			expected bool
		} {
			{"_group", true},
			{"group", false},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %v", tt.gname, tt.expected), func(t *testing.T) {
				actual := Group(tt.gname).IsPrivate()
				assert.Equal(t, tt.expected, actual)
			})
		}
	})



	const (
		positive = "positive"
		negative = "negative"
	)

	type ValidateTestType struct {
		gname string
		expected string
	}

	t.Run("Validate", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group", positive},
			{"group", positive},
			{"g roup", negative},
			{"g~roup", negative},
			{"", negative},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gname, tt.expected), func(t *testing.T) {
				err := Group(tt.gname).Validate()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})

	t.Run("ValidatePublic", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group", negative},
			{"group", positive},
			{"g roup", negative},
			{"g~roup", negative},
			{"", negative},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gname, tt.expected), func(t *testing.T) {
				err := Group(tt.gname).ValidatePublic()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})

	t.Run("ValidateEmptyIsSafe", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group", positive},
			{"group", positive},
			{"g roup", negative},
			{"g~roup", negative},
			{"", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gname, tt.expected), func(t *testing.T) {
				err := Group(tt.gname).ValidateEmptyIsSafe()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})

	t.Run("ValidatePublicEmptyIsSafe", func(t *testing.T) {
		tests := []ValidateTestType {
			{"_group", negative},
			{"group", positive},
			{"g roup", negative},
			{"g~roup", negative},
			{"", positive},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("'%s' is %s", tt.gname, tt.expected), func(t *testing.T) {
				err := Group(tt.gname).ValidatePublicEmptyIsSafe()
				if tt.expected == positive {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err)
				}
				
			})
		}
	})


	t.Run("Match", func(t *testing.T) {
		tests := []struct {
			gname string
			arg string
			expected bool
		} {
			{"_group", "_group", true},
			{"group", "group", true},
			{"group", "", true},
			{"_", "_", true},
			{"_group", "group", false},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Matching '%s' and '%s' is %v", tt.gname, tt.arg, tt.expected), func(t *testing.T) {
				actual := Group(tt.gname).Match(Group(tt.arg))
				assert.Equal(t, tt.expected, actual)
			})
		}
	})
}
