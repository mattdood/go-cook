package cli

// Name is always the first arg, use to discover
// command to run. Flags are the rest
type CommandArgs struct {
    Name string
    Args []string
}

// Pass args to parser then run the
// appropriate command, return exit call
// of the given command
func Run(args []string) int {
    command := CommandArgs{
        Name: args[1],
        Args: args[1:],
    }
    return ParseAndRun(command) // placeholder
}

