package read

import (
	"mdtk/taskset/path"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Markdown_ReadTaskfile(t *testing.T) {

	t.Run("GetTaskfileBlockPath", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			parr, err := Markdown(md_sample).GetTaskfileBlockPath()

			if assert.NoError(t, err) {
				_, ok := parr[path.Path("go/test/file.md")]
				assert.True(t, ok)

				_, ok = parr[path.Path("/markdown/read/data.md")]
				assert.True(t, ok)
			}
		})

		t.Run("negative (error handling of Markdown.ExtractCode)", func(t *testing.T) {})
	})
}
