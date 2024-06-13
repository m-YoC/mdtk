package base

import (
	"fmt"
	"os"
)

const NameReg = "[a-zA-Z_][\\w.-]*"

func GetWorkingDir() (string, error) {
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return p, nil
}

func PairFirst[T, U any](t T, u U) T {
	return t
}

var finalize []func()

func AddFinalize(f func()) {
	finalize = append(finalize, f)
}

func MdtkExit(ecode int) {
	for _, f := range finalize {
		f()
	}
	os.Exit(ecode)
}

func Exit1_IfHasError(err error) {
	if err != nil {
		fmt.Print(err)
		MdtkExit(1)
	}
}

