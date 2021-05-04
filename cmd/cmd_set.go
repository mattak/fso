package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	SetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"s"},
		Short:   "Set data",
		Long:    "Set data",
		Example: `  echo '{"symbols":["a"]}' | fso set projectId collection1 document1
`,
		Run: runCommandSet,
	}
)

func runCommandSet(cmd *cobra.Command, args []string) {
	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		log.Fatal("Please set environment variable: GOOGLE_APPLICATION_CREDENTIALS")
	}

	if len(args) < 3 {
		log.Fatal("write requires 3 arguments: projectId, document, collection")
	}

	projectId := args[0]
	collection := args[1]
	doc := args[2]

	SetChildCollection(projectId, collection, doc)
}
