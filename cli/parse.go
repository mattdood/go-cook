package cli

import (
    "flag"

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
    return cc.fs.Parse(args)
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
    cmd.ParseFlags(command.args)

    return cmd.Run()
}

