package cli

import (
    "errors"
    "flag"
    "fmt"
    "strings"

    "github.com/mattdood/go-cook/run"
)

// Custom array implementation of flag variable
type arrayFlag []string

func (fs *arrayFlag) String() string {
    return strings.Join(*fs, " ")
}

// Custom Usage() override for flags with no arguments
func Usage(name string) {
    fmt.Println("Wraps `git %d`, takes no arguments. Only operates on `cook` directory.", name)
}

// Accepts space separated list of values
func (fs *arrayFlag) Set(value string) error {
    for _, file := range strings.Split(value, " ") {
        *fs = append(*fs, file)
    }
    return nil
}

// Command types, each is required to have a FlagSet
type CreateCommand struct {
    fs *flag.FlagSet
    title string
    template string
}

func NewCreateCommand() *CreateCommand {
    cc := &CreateCommand{
        fs: flag.NewFlagSet("create", flag.ContinueOnError),
    }
    cc.fs.StringVar(&cc.title, "title", "", "Title for the template")
    cc.fs.StringVar(&cc.template, "template", "", "Template type (recipe or tip)")

    return cc
}

func (cc *CreateCommand) ParseFlags(args []string) error {
    err := cc.fs.Parse(args)

    if len(cc.title) == 0 && err != flag.ErrHelp {
        return errors.New("Length of -title flag must be >0 characters")
    }

    if len(cc.template) == 0 && err != flag.ErrHelp {
        return errors.New("Length of -template flag must be >0 characters (select a template)")
    }

    return err
}

func (cc *CreateCommand) Run() int {
    run.Create(cc.title, cc.template)
    return 0
}

// Git command wrapper for adding
// files to be tracked
// `git add <arg>`
type AddCommand struct {
    fs *flag.FlagSet
    files arrayFlag
}

func NewAddCommand() *AddCommand {
    ac := &AddCommand{
        fs: flag.NewFlagSet("add", flag.ContinueOnError),
    }
    ac.fs.Var(&ac.files, "file", "Files to add to git tracking. Space separated list.")

    return ac
}

func (ac *AddCommand) ParseFlags(args []string) error {
    err := ac.fs.Parse(args)

    if len(ac.files) == 0 && err != flag.ErrHelp {
        return errors.New("Length of the -files flag must be >0 characters")
    }

    return err
}

func (ac *AddCommand) Run() int {
    run.Add(ac.files)
    return 0
}

// Git command wrapper for repo init
type InitCommand struct {
    fs *flag.FlagSet
}

func NewInitCommand() *InitCommand {
    ic := &InitCommand{
        fs: flag.NewFlagSet("init", flag.ContinueOnError),
    }
    ic.fs.Usage = Usage(ic.fs.Name)

    return ic
}

func (ic *InitCommand) ParseFlags(args []string) error {
    err := ic.fs.Parse(args)

    if len(args) >= 0 && err != flag.ErrHelp {
        return errors.New("This command takes no arguments")
    }

    return err
}

func (ic *InitCommand) Run() int {
    run.Init()
    return 0
}

// Git command wrapper for repo push
type PushCommand struct {
    fs *flag.FlagSet
}

func NewPushCommand() *PushCommand {
    pc := &PushCommand{
        fs: flag.NewFlagSet("push", flag.ContinueOnError),
    }

    return pc
}

func (pc *PushCommand) ParseFlags(args []string) error {
    err := pc.fs.Parse(args)

    if len(args) >= 0 && err != flag.ErrHelp {
        return errors.New("This command takes no arguments")
    }

    return err
}

func (pc *PushCommand) Run() int {
    run.Push()
    return 0
}

// Git command wrapper for repo pull
type PullCommand struct {
    fs *flag.FlagSet
}

func NewPullCommand() *PullCommand {
    pc := &PullCommand{
        fs: flag.NewFlagSet("pull", flag.ContinueOnError),
    }

    return pc
}

func (pc *PullCommand) ParseFlags(args []string) error {
    err := pc.fs.Parse(args)

    if len(args) >= 0 && err != flag.ErrHelp {
        return errors.New("This command takes no arguments")
    }

    return err
}

func (pc *PullCommand) Run() int {
    run.Pull()
    return 0
}

// TODO:
// - Figure out git subcommands and a choice constant
// - Determine how sub commands are typically parsed out
// - Wrap git subcommands for add, commit, push, pull
type GitCommand struct {
    fs *flag.FlagSet
    subcommand string
}

func NewGitCommand() *GitCommand {
    gc := &GitCommand{
        fs: flag.NewFlagSet("git", flag.ContinueOnError),
    }
    gc.fs.StringVar(&gc.subcommand, "subcommand", "", "Name of the Git subcommand to run")

    return gc
}

func (gc *GitCommand) ParseFlags(args []string) error {
    err := gc.fs.Parse(args)

    if len(gc.subcommand) == 0 && err != flag.ErrHelp {
        return errors.New("Length of the -subcommand flag must be >0 characters")
    }

    return err
}

func (gc *GitCommand) Run() int {
    run.Git(gc.subcommand)
    return 0
}

// Runner interface that passes all
// command functions
type Runner interface {
    ParseFlags([]string) error
    Run() int
}

func ParseAndRun(command CommandArgs) int {

    // Register commands
    cmds := map[string]Runner{
        "create": NewCreateCommand(),
        "add": NewAddCommand(),
        "init": NewInitCommand(),
        "pull": NewPullCommand(),
        "push": NewPushCommand(),
    }

    // Determine cmd that was passed, init,
    // then run
    cmd := cmds[command.name]
    err := cmd.ParseFlags(command.args)

    switch {
    // Usage information for flags is enabled by default
    // if we pass on the `flag.ErrHelp` during arg parsing
    case err == flag.ErrHelp:
        return 0
    case err != nil:
        fmt.Println(err.Error())
        return 2
    }

    return cmd.Run()
}

