package read

import (
    "fmt"
    "os"
	"path/filepath"
	"mdtk/taskset/path"
)

func ReadFile(path path.Path) Markdown {
    b, err := os.ReadFile(path.GetFileAbsPath().FromSlash().String())
    if err != nil {
        fmt.Println("File could not be read.")
        fmt.Println(path)
        os.Exit(1)
    }
    return Markdown(string(b))
}

func SearchTaskfile() path.Path {
	// First: Taskfile.md
	_, err := os.Stat("Taskfile.md")
    if err == nil {
		return "Taskfile.md"
	}

	// Second: *.task.md
	wd, err := path.GetWorkingDir[string]()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	pattern := filepath.Join(wd, "*.taskrun.md")
    files, err := filepath.Glob(pattern)
    if err == nil && len(files) == 1 {
        return path.Path(filepath.Base(files[0]))
    }

	fmt.Println("* No Taskfile found in current directory.")
	fmt.Println("The name of Taskfile is...")
	fmt.Println("  1. Set the filename with the -f option when executing command.")
	fmt.Println("  2. Name the file 'Taskfile.md'.")
	fmt.Println("  3. Name the file '*.taskrun.md'.")
	fmt.Println("Do one of the above. Younger number has priority.")
	fmt.Println("Note, if 'Taskfile.md' does not exist, can create only \n 1 '*.taskrun.md' file (because it cannot determine which file to read).")
	os.Exit(1)

	return ""
}
