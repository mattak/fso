package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	CreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create data",
		Long:    "create data",
		Example: `  echo '{"symbols":["a"]}' | fso create projectId collection1 document1
`,
		Run: runCommandCreate,
	}
)

func runCommandCreate(cmd *cobra.Command, args []string) {
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

	CreateChildCollection(projectId, collection, doc)
}
