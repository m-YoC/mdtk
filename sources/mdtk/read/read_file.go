package read

import (
    "fmt"
    "os"
	"path/filepath"
	"mdtk/path"
)

func ReadFile(path path.Path) Markdown {
    b, err := os.ReadFile(string(path))
    if err != nil {
        fmt.Println("File could not be read.")
        fmt.Println(path)
        os.Exit(1)
    }
    return Markdown(string(b))
}

func GetWorkingDir() string {
	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func SearchTaskfile() path.Path {
	// First: Taskfile.md
	_, err := os.Stat("Taskfile.md")
    if err == nil {
		return "Taskfile.md"
	}

	 // Second: *.task.md
	pattern := filepath.Join(GetWorkingDir(), "*.taskrun.md")
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
