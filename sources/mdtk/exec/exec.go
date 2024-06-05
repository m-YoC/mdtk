package exec

import (
	"fmt"
	"os"
	"strings"
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

	out, err := exec.Command(GetShell(), "-c", execcode).CombinedOutput()
	if quiet_mode {
		out = nil	
	}
	if err != nil {	
		PrintExecError(out)
		os.Exit(1)
	}

	// 実行したコマンドの結果を出力
	fmt.Print(string(out))
}

func PrintExecError(out []byte) {
	errtext := "mdtk: Command exec error."
	fmt.Print(string(out))
	if !strings.Contains(string(out), errtext) {
		fmt.Println(errtext, "Error command was run in", os.Args)
	}
}
