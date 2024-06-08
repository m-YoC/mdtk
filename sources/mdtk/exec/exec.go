package exec

import (
	"fmt"
	"os"
	"os/exec"
	"mdtk/config"
)

func GetShell() string {
	return config.Config.Shell
}

func GetShHead() string {
	return config.Config.ScriptHeadSet
}

func Run(code string, quiet_mode bool) {
	cmd := exec.Command(GetShell(), "-c", GetShHead() + "\n" + code)

	if !quiet_mode {
		cmd.Stdout = os.Stdout
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {	
		errtext := "mdtk: Command exec error."
		fmt.Println(errtext, "Error command was run in", os.Args)
		os.Exit(1)
	}
}

