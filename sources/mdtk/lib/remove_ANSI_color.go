package lib

import(
	"regexp"
)

const ansi_color_reg = `\x1b\[[0-9;]*m`
var ansi_color_rex = regexp.MustCompile(ansi_color_reg)

func RemoveANSIColor(str string) string {
	return ansi_color_rex.ReplaceAllString(str, "")
}
