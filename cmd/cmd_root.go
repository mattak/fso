package cmd

import (
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "crud",
		Short: "firebase get/set operator",
		Long:  "firebase get/set operator",
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(GetCmd)
	RootCmd.AddCommand(SetCmd)
	RootCmd.AddCommand(CreateCmd)
	RootCmd.AddCommand(DeleteCmd)
}
