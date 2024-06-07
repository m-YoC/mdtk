package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_setConfig(t *testing.T) {
	assert.Equal(t, "/bin/bash", Config.Shell)

	setConfig([]string{"shell: /bin/sh"})
	assert.Equal(t, "/bin/sh", Config.Shell)

	assert.Error(t, setConfig([]string{"hoge: /bin/bash"}))
	assert.Error(t, setConfig([]string{"pager: less 'a"}))
	assert.Error(t, setConfig([]string{"pager_min_limit: PI"}))
	assert.Error(t, setConfig([]string{"pager_min_limit: 3.14"}))
	assert.Error(t, setConfig([]string{"pager_min_limit: -20"}))
}
