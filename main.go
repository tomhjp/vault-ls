package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tomhjp/vault-ls/internal/cmd"
)

func main() {
	err := realMain()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	if len(os.Args) < 2 {
		return errors.New("Must pass at least 1 argument")
	}

	switch os.Args[1] {
	case "serve":
		c := &cmd.ServeCommand{
			Version: VersionString(),
		}
		return c.Run(os.Args[1:])
	case "version":
		c := &cmd.VersionCommand{
			Version: VersionString(),
		}
		return c.Run(os.Args[1:])
	default:
		return fmt.Errorf("unsupported command: %s", os.Args[1])
	}
}
