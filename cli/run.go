package cli

// Name is always the first arg, use to discover
// command to run. Flags are the rest
type CommandArgs struct {
    name string
    args []string
}

// Pass args to parser then run the
// appropriate command, return exit call
// of the given command
func Run(args []string) int {
    if len(args) < 1 {
        return 1
    }

    command := CommandArgs{
        name: args[1],
        args: args[2:],
    }
    return ParseAndRun(command) // placeholder
}

