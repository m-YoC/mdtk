package read

import (
	"mdtk/group"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Markdown_ReadTaskConfigGroupOrder(t *testing.T) {

	t.Run("GetTaskConfigGroupOrder", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			parr, err := Markdown(md_sample).GetTaskConfigGroupOrder()

			if assert.NoError(t, err) {
				v, ok := parr[group.Group("hello")]
				assert.True(t, ok)
				assert.Equal(t, int64(1), v)

				v, ok = parr[group.Group("world")]
				assert.True(t, ok)
				assert.Equal(t, int64(-2), v)
			}
		})

		t.Run("negative: value is not integer", func(t *testing.T) {
			parr, err := Markdown(md_sample).GetTaskConfigGroupOrder()

			if assert.NoError(t, err) {
				_, ok := parr[group.Group("nothing")]
				assert.False(t, ok)
			}
		})
	})
}
