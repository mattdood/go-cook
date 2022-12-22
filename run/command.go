package run

import (
    "fmt"
)

func Create(title string, template string) {

    fmt.Println("Hello from the command.go file: ", title)
}

func Add(files []string) {
    fmt.Println("Hello from the (git) add command: ", files)
}

func Init() {
    fmt.Println("Hello from the init command")
}

func Push() {
    fmt.Println("Hello from the push command")
}

func Pull() {
    fmt.Println("Hello from the pull command")
}
