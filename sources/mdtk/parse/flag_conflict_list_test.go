package parse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_FlagConflictList(t *testing.T) {
	t.Run("Conflict", func(t *testing.T) {
		var list FlagConflictList
		assert.Equal(t, 0, len(list.List))

		fs := []string{"-a", "-b", "-c"}
		list.Conflict(fs...)
		assert.Equal(t, 1, len(list.List))
		assert.Equal(t, fs, list.List[0])
	})

	t.Run("Check", func(t *testing.T) {
		{
			var list FlagConflictList
			fs := []string{"-a", "-b", "-c"}
			list.Conflict(fs...)
	
			var flags Flag
			flags.Set("-a")
			flags.Set("-b")
			flags.Set("-c")
			flags[flags.GetIndex("-a")].Exist = true
			flags[flags.GetIndex("-c")].Exist = true
			assert.Error(t, list.Check(flags))
		}

		{
			var list FlagConflictList
			fs := []string{"-a", "-b", "-c"}
			list.Conflict(fs...)
	
			var flags Flag
			flags.Set("-a")
			flags.Set("-b")
			flags.Set("-c")
			flags[flags.GetIndex("-c")].Exist = true
			assert.NoError(t, list.Check(flags))
		}
	})
}
