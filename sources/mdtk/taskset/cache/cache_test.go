package cache

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_toCacheName(t *testing.T) {
	assert.Equal(t, "Taskfile.md.cache", toCacheName("Taskfile.md"))
}
