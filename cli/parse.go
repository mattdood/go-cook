package cli

import (
    "errors"
    "flag"
    "fmt"
    "strings"

    "github.com/mattdood/go-cook/run"
)

// Command types, each is required to have a FlagSet
type CreateCommand struct {
    fs *flag.FlagSet
    title string
}

func NewCreateCommand() *CreateCommand {
    cc := &CreateCommand{
        fs: flag.NewFlagSet("create", flag.ContinueOnError),
    }
    cc.fs.StringVar(&cc.title, "title", "", "Title for the template")

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
    run.Create(cc.title)
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

