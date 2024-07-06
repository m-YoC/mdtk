package lib

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func CmdTest(cmd string) string {
	c := strings.Fields(cmd)
	res, err := exec.Command(c[0], c[1:]...).CombinedOutput()
	if err != nil {
		return fmt.Sprint(string(res), err)
	} else {
		return string(res)
	}
}

func AssertStringContains(t *testing.T, substr, str string) {
	res := strings.Contains(str, substr)
	assert.True(t, res)
	if !res {
		t.Logf("Actual             : %q", str)
		t.Logf("Expected to contain: %q", substr)
	}
}

