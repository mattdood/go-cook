package run

import (
    "fmt"
    "os/exec"
    "strings"
)

const (
    TemplatesBaseDirectory string = "templates/"
    OutputBaseDirectory string = "~/cook/"
)

func Create(title string, template string) {

    fmt.Println("Hello from the command.go file: ", title)
}

// Git command wrapper for `git add`
func Add(files []string) {
    out, err := exec.Command("git", "add", strings.Join(files, " ")).Output()
    if err != nil {
        fmt.Println("`git add` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Println(output)
}

// Git command wrapper for `git init`
func Init() {
    out, err := exec.Command("git", "init", OutputBaseDirectory).Output()
    if err != nil {
        fmt.Println("`git init` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Println(output)
}

// Git command wrapper for `git push`
func Push() {
    out, err := exec.Command("git", "push", OutputBaseDirectory).Output()
    if err != nil {
        fmt.Println("`git push` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Println(output)
}

// Git command wrapper for `git pull`
func Pull() {
    out, err := exec.Command("git", "pull", OutputBaseDirectory).Output()
    if err != nil {
        fmt.Println("`git push` exited abnormally")
        fmt.Println(err)
    }
    output := string(out[:])
    fmt.Println(output)
}
