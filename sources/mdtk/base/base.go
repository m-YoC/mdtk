package base

import (
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



