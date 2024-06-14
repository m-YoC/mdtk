package read

import (
	"mdtk/taskset/path"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Markdown_ReadTaskfile(t *testing.T) {

	t.Run("GetTaskfileBlockPath", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			parr, err := Markdown(md_sample).GetTaskfileBlockPath("/hello/root.md")

			if assert.NoError(t, err) {
				v, ok := parr[path.Path("/hello/go/test/file.md")]
				assert.True(t, ok)
				assert.False(t, v)

				v, ok = parr[path.Path("/hello/markdown/read/data.md")]
				assert.True(t, ok)
				assert.False(t, v)

				v, ok = parr[path.Path("/hello/root.md")]
				assert.True(t, ok)
				assert.True(t, v)
			}
		})

		t.Run("negative (error handling of Markdown.ExtractCode)", func(t *testing.T) {})
	})
}
