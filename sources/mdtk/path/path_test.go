package path

import (
	"fmt"
	"regexp"
	"testing"
	"github.com/stretchr/testify/assert"
	"os/user"
)

func Test_Path(t *testing.T) {
	usr, _ := user.Current()

	t.Run("homeDirToAbs", func(t *testing.T) {
		tests := []struct {
			actual string
			expected string
		} {
			{"~/", usr.HomeDir + "/"},
			{"/hello/world", "/hello/world"},
			{"hello/world", "hello/world"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s equals %s", tt.actual, tt.expected), func(t *testing.T) {
				actual := Path(tt.actual).homeDirToAbs()
				expected := Path(tt.expected)
			
				assert.Equal(t, expected, actual)
				
			})
		}
	})
	
	t.Run("GetFileAbsPath", func(t *testing.T) {
		tests := []struct {
			actual string
			expected string
		} {
			{"~/hello", "^" + usr.HomeDir + "/hello$"},
			{"/hello/world", "^/hello/world$"},
			{"hello/world", "^/.+/path/hello/world$"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s equals %s", tt.actual, tt.expected), func(t *testing.T) {
				actual := Path(tt.actual).GetFileAbsPath()
			
				assert.Regexp(t, regexp.MustCompile(tt.expected), actual)
				
			})
		}
	})

	t.Run("Dir", func(t *testing.T) {
		tests := []struct {
			actual string
			expected string
		} {
			{"/hello/world", "/hello"},
			{"/hello/world/", "/hello/world"},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s equals %s", tt.actual, tt.expected), func(t *testing.T) {
				actual := Path(tt.actual).Dir()
				expected := Path(tt.expected)
			
				assert.Equal(t, expected, actual)
				
			})
		}
	})
}
