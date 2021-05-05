package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	GetCmd = &cobra.Command{
		Use:   "get",
		Aliases: []string{"g"},
		Short: "Get data",
		Long:  "Get data",
		Example: `  fso get projectId collection1 document1
`,
		Run: runCommandGet,
	}
	isMeta = false
)

func init() {
	GetCmd.PersistentFlags().BoolVarP(&isMeta, "meta", "m", false, "read meta data")
}

func runCommandGet(cmd *cobra.Command, args []string) {
	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		log.Fatal("Please set environment variable: GOOGLE_APPLICATION_CREDENTIALS")
	}

	if len(args) == 3 {
		projectId := args[0]
		collection := args[1]
		doc := args[2]
		GetChildCollections(projectId, collection, doc, isMeta)
	} else {
		log.Fatal("read requires 3 arguments: projectId, document, collection")
	}
}
