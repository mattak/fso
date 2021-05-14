package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	VersionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show version",
		Long:    "Show version",
		Example: `  fso version
`,
		Run: runCommandVersion,
	}
)

func runCommandVersion(cmd *cobra.Command, args []string) {
	fmt.Println(VERSION)
}
