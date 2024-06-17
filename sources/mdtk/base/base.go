package base

import (
	"os"
	"strings"
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


func ToLF(str string) string {
	return strings.Replace(str, "\r\n", "\n", -1)
}

func ToCRLF(str string) string {
	return strings.Replace(str, "\n", "\r\n", -1)
}

