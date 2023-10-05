package main

import (
	"github.com/aldoborrero/ethw/internal/cmd"
	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cmd.Cli)

	// Configure logging
	cmd.Cli.Log.ConfigureLog()

	// Run the appropiate sub-command
	ctx.FatalIfErrorf(ctx.Run())
}
