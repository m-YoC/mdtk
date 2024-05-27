package exec

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

var shname string

func init() {
	shname = getShell()
}

func Shname() string {
	return shname
}

func getShell() string {
	sh := os.Getenv("SHELL")
	if sh == "" {
		sh = "sh"
	}

	return sh
}

func ToExecCode(code string, eos string) string {
	shhead := "set -euo pipefail\n"
	return "cat - << '" + eos + "' | " + shname + "\n" + shhead + code + "\n" + eos
}

func Run(code string) {
	execcode := ToExecCode(code, "EOS")
	// fmt.Println(execcode)

	out, err := exec.Command(Shname(), "-c", execcode).CombinedOutput()
	if err != nil {
		PrintExecError(out)
		os.Exit(1)
	}

	// 実行したコマンドの結果を出力
	fmt.Printf("%s", string(out))
}

func PrintExecError(out []byte) {
	errtext := "mdtk: Command exec error."
	fmt.Print(string(out))
	if !strings.Contains(string(out), errtext) {
		fmt.Println(errtext, "Error command was run in", os.Args)
	}
}
