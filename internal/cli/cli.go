package cli

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	JSON JSONCmd `cmd:"" help:"Compare JSON files."`
	YAML YAMLCmd `cmd:"" help:"Compare YAML files."`
}

func NewCLI() *CLI {
	return &CLI{}
}

// ParseAndRun is a helper to parse CLI arguments and execute the corresponding command.
func (cli *CLI) ParseAndRun(args []string) error {
	ctx := kong.Parse(cli, kong.Name("unordiff"), kong.Description("A tool for comparing JSON and YAML files without considering order."), kong.UsageOnError())
	return ctx.Run()
}
