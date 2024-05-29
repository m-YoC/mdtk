package help

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PagerOutput(s string) {
	pager := os.Getenv("PAGER")
	if pager == "" { pager = "less" }

    cmd := exec.Command(pager)
    cmd.Stdin = strings.NewReader(s)
    cmd.Stdout = os.Stdout

    err := cmd.Run()
    if err != nil {
        fmt.Println(s)
    }
}
