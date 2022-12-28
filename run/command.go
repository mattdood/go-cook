package run

import (
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path"
    "strings"
)

var (
    OutputBaseDirectory string = getUserHome()
    OutputDirectory string = path.Join(OutputBaseDirectory, InstallBaseDirectory)
)

const (
    InstallBaseDirectory string = "cook/"
    TemplatesBaseDirectory string = "templates/"
)

func getUserHome() string {
    homedir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
    }
    return homedir
}

func Create(title string, template string) {
    fmt.Println("Hello from the command.go file: ", title)
}

// Git command wrapper for `git add`
// TODO:
//   * Should we check in the file path of each file for the `homedir/cook/` as a prefix then prepend?
func Add(files []string) {
    out, err := exec.Command("git", "add", strings.Join(files, " ")).Output()
    if err != nil {
        fmt.Println("`git add` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Print(output)
}

// Git command wrapper for `git init`
func Init() {
    if _, err := os.Stat(OutputDirectory); errors.Is(err, os.ErrNotExist) {
        os.Mkdir(OutputDirectory, 0775)
    }

    out, err := exec.Command("git", "init", OutputDirectory).Output()
    if err != nil {
        fmt.Println("`git init` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Print(output)
}

// Git command wrapper for `git push`
func Push() {
    out, err := exec.Command("git", "push", OutputDirectory).Output()
    if err != nil {
        fmt.Println("`git push` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Print(output)
}

// Git command wrapper for `git pull`
func Pull() {
    out, err := exec.Command("git", "pull", OutputDirectory).Output()
    if err != nil {
        fmt.Println("`git push` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Print(output)
}
