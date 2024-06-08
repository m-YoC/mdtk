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

func ToExecCode(code string, eos string) string {
	return "cat - << '" + eos + "' | " + GetShell() + "\n" + GetShHead() + "\n" + code + "\n" + eos
}

func Run(code string, quiet_mode bool) {
	execcode := ToExecCode(code, "EOS")
	// fmt.Println(execcode)

	cmd := exec.Command(GetShell(), "-c", execcode)
	if !quiet_mode {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {	
		errtext := "mdtk: Command exec error."
		fmt.Println(errtext, "Error command was run in", os.Args)
		os.Exit(1)
	}

	// 実行したコマンドの結果を出力
	// fmt.Print(string(out))
}

