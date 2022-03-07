package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type VersionOutput struct {
	Version string `json:"version"`

	*BuildInfo
}

type VersionCommand struct {
	Version string

	jsonOutput bool
}

type BuildInfo struct {
	GoVersion string `json:"go,omitempty"`
	GoOS      string `json:"os,omitempty"`
	GoArch    string `json:"arch,omitempty"`
	Compiler  string `json:"compiler,omitempty"`
}

func (c *VersionCommand) flags() *flag.FlagSet {
	fs := defaultFlagSet("version")

	fs.BoolVar(&c.jsonOutput, "json", false, "output the version information as a JSON object")

	fs.Usage = func() {
		_, _ = fmt.Fprint(os.Stdout, c.Help())
	}

	return fs
}

func (c *VersionCommand) Run(args []string) error {
	f := c.flags()
	if err := f.Parse(args); err != nil {
		return fmt.Errorf("Error parsing command-line flags: %s", err)
	}

	output := VersionOutput{
		Version: c.Version,
		BuildInfo: &BuildInfo{
			GoVersion: runtime.Version(),
			GoOS:      runtime.GOOS,
			GoArch:    runtime.GOARCH,
			Compiler:  runtime.Compiler,
		},
	}

	if c.jsonOutput {
		jsonOutput, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return fmt.Errorf("\nError marshalling JSON: %s", err)
		}
		_, _ = fmt.Fprint(os.Stdout, jsonOutput)
		return nil
	}

	ver := fmt.Sprintf("%s\nplatform: %s/%s\ngo: %s\ncompiler: %s", c.Version, output.GoOS, output.GoArch, output.GoVersion, output.Compiler)
	_, _ = fmt.Fprint(os.Stdout, ver)
	return nil
}

func (c *VersionCommand) Help() string {
	helpText := `
Usage: vault-ls version [-json]

` + c.Synopsis() + "\n\n" + helpForFlags(c.flags())

	return strings.TrimSpace(helpText)
}

func (c *VersionCommand) Synopsis() string {
	return "Displays the version of the language server"
}
