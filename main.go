package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tomhjp/vault-ls/internal/cmd"
)

// func main() {
// 	err := realMain()
// 	if err != nil {
// 		_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
// 		os.Exit(1)
// 	}
// }

func main() {
	c := &cmd.WASMCommand{
		Version: VersionString(),
	}
	if err := c.Run(os.Args); err != nil {
		_, _ = fmt.Println("error running wasm command", err)
		os.Exit(1)
	}
}

type command interface {
	Run(args []string) error
}

func realMain() error {
	if len(os.Args) < 2 {
		return errors.New("Must pass at least 1 argument")
	}

	var c command
	switch os.Args[1] {
	case "serve":
		c = &cmd.ServeCommand{
			Version: VersionString(),
		}
	case "version":
		c = &cmd.VersionCommand{
			Version: VersionString(),
		}
	case "wasm":
		c = &cmd.WASMCommand{
			Version: VersionString(),
		}
	default:
		return fmt.Errorf("unsupported command: %s", os.Args[1])
	}

	return c.Run(os.Args[1:])
}
