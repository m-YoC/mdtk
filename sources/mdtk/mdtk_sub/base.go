package sub

import (
	"os"
)

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
