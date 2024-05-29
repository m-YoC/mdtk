package help

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PagerOutput(s string, active_num uint) {
	s = strings.TrimRight(s, "\n") + "\n\n"
	count := strings.Count(s, "\n")
	if count < int(active_num) {
		fmt.Print(s)
		return
	}

	pager := os.Getenv("PAGER")
	var cmd *exec.Cmd
	if pager == "" { 
		cmd = exec.Command("less", "-R")
	} else if strings.HasSuffix(pager, "less") {
		cmd = exec.Command(pager, "-R")
	} else {
		cmd = exec.Command(pager)
	}
    
    cmd.Stdin = strings.NewReader(s)
    cmd.Stdout = os.Stdout

    err := cmd.Run()
    if err != nil {
        fmt.Print(s)
    }
}
