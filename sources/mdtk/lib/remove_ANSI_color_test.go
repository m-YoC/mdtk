package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const data = "\x1b[39mdescription\x1b[90m (+5)\x1b[0m"

func Test_RemoveANSIColor(t *testing.T) {
	
	assert.Equal(t, "description (+5)", RemoveANSIColor(data))
}
