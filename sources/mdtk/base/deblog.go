package base

import (
	"fmt"
	"github.com/gookit/color"
)

var use_debug_log bool = false
var layer_func func(int)string

func UseDebugLog(f func(int)string) {
	use_debug_log = true
	layer_func = f
}

func DebugLogNoLayer(str string) {
	if use_debug_log {
		fmt.Print(str)
	}
}

func DebugLog(layer int, str string) {
	if use_debug_log {
		fmt.Print(layer_func(layer) + " ")
		fmt.Print(str)
	}
}

func DebugLogGreen(layer int, str string) {
	if use_debug_log {
		fmt.Print(layer_func(layer) + " ")
		color.Green.Print(str)
	}
}

func DebugLogMagenta(layer int, str string) {
	if use_debug_log {
		fmt.Print(layer_func(layer) + " ")
		color.Magenta.Print(str)
	}
}

func DebugLogGray(layer int, str string) {
	if use_debug_log {
		fmt.Print(layer_func(layer) + " ")
		color.Gray.Print(str)
	}
}


func DebugLogExit() {
	if use_debug_log {
		color.Yellow.Println("Task is not run when debug mode.")
		MdtkExit(0)
	}
}

