package run

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"
)

var (
	OutputBaseDirectory = getUserHome()
	OutputDirectory     = path.Join(OutputBaseDirectory, InstallBaseDirectory)
)

const (
	InstallBaseDirectory string      = "cook/"
	FilePermission       fs.FileMode = 00775
)

func getUserHome() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	return homedir
}

type TemplateData struct {
	Title     string
	Category  string
	Timestamp string
	Tags      []string
}

func NewTemplateData(title string, category string, tags []string, timestamp string) *TemplateData {
	return &TemplateData{
		Title:     title,
		Category:  category,
		Tags:      tags,
		Timestamp: timestamp,
	}
}

// Create a new file based on the template data input via the CLI
func Create(title string, category string, tags []string, templateType string) int {
	// Find template file to load
	var tmpl *template.Template

	switch {
	case templateType == "recipe":
		// Seems to need the "run/" due to it executing in main.go one level up?
		var err error

		tmpl, err = template.ParseFiles("run/recipe.tmpl")
		if err != nil {
			fmt.Println("Error parsing `recipe.tmpl` template file: ", err)
			return 1
		}
	case templateType == "tip":
		var err error

		tmpl, err = template.ParseFiles("run/tip.tmpl")
		if err != nil {
			fmt.Println("Error parsing `tip.tmpl` template file: ", err)
			return 1
		}
	}

	// Path interpolation
	timestamp := time.Now()

	filename := fmt.Sprintf("%s-%s.yml", strings.Replace(title, " ", "-", -1), timestamp.Format("20060102150405"))

	outputPath := path.Join(
		OutputDirectory,
		fmt.Sprintf("%ss", templateType),
		fmt.Sprintf("%d/%02d/%02d/", timestamp.Year(), timestamp.Month(), timestamp.Day()),
		timestamp.Format("20060102150405"),
		filename,
	)

	// Create all directories to the file
	if err := os.MkdirAll(path.Dir(outputPath), FilePermission); err != nil {
		fmt.Println(err)
		return 1
	}

	// Create file
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	// Load template data
	err = tmpl.Execute(
		file,
		NewTemplateData(title, category, tags, timestamp.Format("20060102150405")),
	)
	if err != nil {
		fmt.Println("Error in template execution: ", err)
		return 1
	}

	file.Close()

	// Open in vim
	cmd := exec.Command("vim", outputPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error in vim execution: ", err)
		return 1
	}

	return 0
}

// Git command wrapper for `git add`
func Add(files []string) {
	out, err := exec.Command("git", "-C", OutputDirectory, "add", strings.Join(files, " ")).Output()
	if err != nil {
		fmt.Println("`git add` exited abnormally")
		fmt.Println(err)
	}

	output := string(out)

	fmt.Print(output)
}

// Git command wrapper for `git init`
func Init() {
	if _, err := os.Stat(OutputDirectory); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(OutputDirectory, FilePermission)
		if err != nil {
			fmt.Println("Failed to create directories")
			fmt.Println(err)
		}
	}

	out, err := exec.Command("git", "-C", OutputDirectory, "init").Output()
	if err != nil {
		fmt.Println("`git init` exited abnormally")
		fmt.Println(err)
	}

	output := string(out)

	fmt.Print(output)
}

// Git command wrapper for `git push`
func Push() {
	out, err := exec.Command("git", "-C", OutputDirectory, "push").Output()
	if err != nil {
		fmt.Println("`git push` exited abnormally")
		fmt.Println(err)
	}

	output := string(out)

	fmt.Print(output)
}

// Git command wrapper for `git pull`
func Pull() {
	out, err := exec.Command("git", "-C", OutputDirectory, "pull").Output()
	if err != nil {
		fmt.Println("`git push` exited abnormally")
		fmt.Println(err)
	}

	output := string(out[:])

	fmt.Print(output)
}
