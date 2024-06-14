package exec

import (
	"fmt"
	"os"
	"os/exec"
	"mdtk/config"
	"mdtk/taskset/path"
)

func GetShell() string {
	return config.Config.Shell[0]
}

func GetShellOpt() []string {
	if len(config.Config.Shell) == 1 { 
		return []string{}
	} else {
		return config.Config.Shell[1:]
	}
}

func GetShHead() string {
	return config.Config.ScriptHeadSet
}

func Run(code string, fdir string, quiet_mode bool, rfd bool) error {
	if rfd {
		prev, err := path.GetWorkingDir[string]()
		if err != nil {
			return err
		}
		defer os.Chdir(prev)
		os.Chdir(fdir)
	}
	
	cmd := exec.Command(GetShell(), append(GetShellOpt(), GetShHead() + "\n" + code)...)

	if !quiet_mode {
		cmd.Stdout = os.Stdout
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {	
		errtext := "mdtk: Command exec error."
		s := fmt.Sprintln(errtext, "Error command was run in", os.Args)
		return fmt.Errorf("%s\n", s)
	}

	return nil
}

