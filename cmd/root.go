package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use:   "fso",
		Short: "firebase get/set operator",
		Long:  "firebase get/set operator",
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(GetCmd)
	RootCmd.AddCommand(SetCmd)
	RootCmd.AddCommand(CreateCmd)
}
