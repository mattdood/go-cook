package run

import (
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path"
    "strings"
    "text/template"
    "time"
)

var (
    OutputBaseDirectory string = getUserHome()
    OutputDirectory string = path.Join(OutputBaseDirectory, InstallBaseDirectory)
)

const (
    InstallBaseDirectory string = "cook/"
)

func getUserHome() string {
    homedir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
    }
    return homedir
}

type TemplateData struct {
    Title string
    Category string
    Tags []string
    Timestamp string
}

func NewTemplateData(title string, category string, tags []string, timestamp string) *TemplateData {
    return &TemplateData{
        Title: title,
        Category: category,
        Tags: tags,
        Timestamp: timestamp,
    }
}

// Create a new file based on the template data input via the CLI
func Create(title string, category string, tags []string, templateType string) int {
    // Find template file to load
    var tmpl *template.Template
    var err error
    switch {
    case templateType == "recipe":
        tmpl, err = template.ParseFiles("run/recipe.tmpl")
        if err != nil {
            fmt.Println("Error parsing `recipe.tmpl` template file: ", err)
            return 1
        }
    case templateType == "tip":
        tmpl, err = template.ParseFiles("run/tip.tmpl")
        if err != nil {
            fmt.Println("Error parsing `tip.tmpl` template file: ", err)
            return 1
        }
    }

    // Path interpolation
    var timestamp time.Time
    timestamp = time.Now()

    var filename string
    filename = fmt.Sprintf("%s-%s.yml", strings.Replace(title, " ", "-", -1), timestamp.Format("20060102150405"))

    var outputPath string
    outputPath = path.Join(
        OutputDirectory,
        templateType,
        fmt.Sprintf("%d/%02d/%02d/", timestamp.Year(), timestamp.Month(), timestamp.Day()),
        timestamp.Format("20060102150405"),
        filename,
    )

    // Create all directories to the file
    if err := os.MkdirAll(path.Dir(outputPath), 0775); err != nil {
        fmt.Println(err)
        return 2
    }

    // Create file
    file, err := os.Create(outputPath)
    if err != nil {
        fmt.Println(err)
        return 2
    }

    // Load template data
    err = tmpl.Execute(
        file,
        NewTemplateData(title, category, tags, timestamp.Format("20060102150405")),
    )
    if err != nil {
        fmt.Println("Error in template execution: ", err)
        return 2
    }
    file.Close()

    // Open in vim
    cmd := exec.Command("vim", outputPath)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error in vim execution: ", err)
        return 2
    }

    return 0
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
