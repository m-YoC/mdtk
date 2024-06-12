package code

import (
	"fmt"
	"mdtk/exec"
)

func (code Code) GetRunnableScript() string {
	h := fmt.Sprintln("#!" + exec.GetShell()) //shebang
	h += fmt.Sprintln(exec.GetShHead())
	h += fmt.Sprintln("")

	return h + string(code.RemoveEmbedDescComment().RemoveEmbedArgsComment())
}

func (code Code) GetRawScript() string {
	return string(code.RemoveEmbedDescComment().RemoveEmbedArgsComment())
}
