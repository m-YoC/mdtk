package help

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"mdtk/config"
)

func PagerOutput(s string, pager_min_row uint) {
	s = strings.TrimRight(s, "\n") + "\n\n"
	count := strings.Count(s, "\n")
	pager := config.Config.Pager

	if count < int(pager_min_row) || len(pager) == 0 {
		fmt.Print(s)
		return
	}

	var cmd *exec.Cmd
	if strings.HasSuffix(pager[0], "less") && len(pager) == 1 {
		cmd = exec.Command(pager[0], "-R")
	} else {
		cmd = exec.Command(pager[0], pager[1:]...)
	}
    
    cmd.Stdin = strings.NewReader(s)
    cmd.Stdout = os.Stdout

    err := cmd.Run()
    if err != nil {
        fmt.Print(s)
    }
}
