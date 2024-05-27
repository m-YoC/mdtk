package read

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Markdown_ReadTaskfile(t *testing.T) {

	t.Run("GetTaskfileBlock", func(t *testing.T) {
		t.Run("positive", func(t *testing.T) {
			parr, err := Markdown(md_sample).GetTaskfileBlockPath()

			if assert.NoError(t, err) {
				assert.Equal(t, "go/test/file.md", string(parr[0]))
				assert.Equal(t, "/markdown/read/data.md", string(parr[1]))
			}
		})

		t.Run("negative (error handling of Markdown.ExtractCode)", func(t *testing.T) {})
	})
}
