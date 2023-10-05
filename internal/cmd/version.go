package cmd

import (
	"fmt"
	"os"

	"github.com/aldoborrero/ethw/internal/build"
	"github.com/alecthomas/kong"
)

type versionFlag bool

func (v versionFlag) Decode(ctx *kong.DecodeContext) error {
	fmt.Printf("%s version %s\n", build.Name, build.Version)
	os.Exit(0)
	return nil
}
