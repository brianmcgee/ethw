package cmd

import (
	"fmt"

	"github.com/aldoborrero/ethw/internal/build"
)

type versionCmd struct{}

func (cmd *versionCmd) Run() error {
	fmt.Printf("%s version %s\n", build.Name, build.Version)
	return nil
}
