package code

import (
	"fmt"
)

func (code Code) GetRunnableScript(shpath string, head string) string {
	h := fmt.Sprintln("#!" + shpath) //shebang
	h += fmt.Sprintln(head)
	h += fmt.Sprintln("")

	return h + code.GetRawScript()
}

func (code Code) GetRawScript() string {
	return string(code.RemoveEmbedDescComment().RemoveEmbedArgsComment())
}
