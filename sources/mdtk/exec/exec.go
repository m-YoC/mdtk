package exec

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
	"mdtk/taskset/path"
)

type LangInterface interface {
	GetCmd(string, bool) (string, []string, func())
}

func GetCommand(first string, other []string, rmf func()) (*exec.Cmd, func()) {
	return exec.Command(first, other...), rmf
}

func Run(lang LangInterface, code string, use_tmpfile_mode bool, quiet_mode bool, rfd bool, fdir string) error {
	if rfd {
		prev, err := path.GetWorkingDir[string]()
		if err != nil {
			return err
		}
		defer os.Chdir(prev)
		os.Chdir(fdir)
	}

	cmd, rmf := GetCommand(lang.GetCmd(code, use_tmpfile_mode))
	defer rmf()
	
	if !quiet_mode {
		cmd.Stdout = os.Stdout
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {	
		errtext := "mdtk: Command exec error."
		s := fmt.Sprintln(errtext, "Error command was run in", os.Args)
		s += fmt.Sprintf("Is [%s] already installed? Can you run it?\n", strings.Fields(cmd.String())[0])
		return fmt.Errorf("%s\n", s)
	}

	return nil
}

