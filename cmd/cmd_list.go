package cmd

import (
	"github.com/mattak/fso/pkg/crud"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	ListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Short:   "List data",
		Long:    "List data",
		Example: `  crud list projectId
  crud list projectId collection1
`,
		Run: runCommandList,
	}
)

func runCommandList(cmd *cobra.Command, args []string) {
	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		log.Fatal("Please set environment variable: GOOGLE_APPLICATION_CREDENTIALS")
	}

	if len(args) == 1 {
		projectId := args[0]
		result := crud.ListUpRootCollections(projectId)
		crud.PrintMapTsv(result)
	} else if len(args) == 2 {
		projectId := args[0]
		collection := args[1]
		result := crud.ListUpDocs(projectId, collection)
		crud.PrintMapTsv(result)
	} else {
		log.Fatal("read requires 1-2 arguments: projectId, document")
	}
}
