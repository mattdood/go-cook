package cli

import (
    "fmt"
)

const usage = `
go-cook [OPTION] [args...]

Options:
    -h, -help, --help display this help

Executes various commands to manipulate cooking templates defined
in YAML format.

Example:
$ go-cook create -title "my-recipe-name"`

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

    // args are:
    // [/tmp/go-build-dir <cmd flag>]
    if len(args) > 1{
        fmt.Println(args)
        for _, arg := range args {
            fmt.Println(arg)
            if arg == "-h" || arg == "-help" || arg == "--help" {
                fmt.Println(usage)
                return 1
            }
        }
    } else {
        fmt.Println("Unknown command. Use -h, -help, or --help to display help")
        return 1
    }

    command := CommandArgs{
        name: args[1],
        args: args[2:],
    }

    return ParseAndRun(command)
}

