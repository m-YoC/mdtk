package base

import (
	"fmt"
	"os"
)

var osExit = os.Exit

var finalize []func()

func AddFinalize(f func()) {
	finalize = append(finalize, f)
}

func MdtkExit(ecode int) {
	for _, f := range finalize {
		f()
	}
	osExit(ecode)
}

func Exit1_IfHasError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		MdtkExit(1)
	}
}

// for Test
// defer NewExit(func(status int) { test_status = status })()
func NewExit(exit func(status int)) func() {
	old := osExit
	osExit = exit
	return func() { osExit = old }
}
