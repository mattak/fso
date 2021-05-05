package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	DeleteCmd = &cobra.Command{
		Use:   "delete",
		Aliases: []string{"d"},
		Short: "Delete data",
		Long:  "Delete data",
		Example: `  fso delete projectId collection1 document1
`,
		Run: runCommandDelete,
	}
)

func runCommandDelete(cmd *cobra.Command, args []string) {
	_, exists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if !exists {
		log.Fatal("Please set environment variable: GOOGLE_APPLICATION_CREDENTIALS")
	}

	if len(args) == 3 {
		projectId := args[0]
		collection := args[1]
		doc := args[2]
		DeleteChildCollection(projectId, collection, doc)
	} else {
		log.Fatal("read requires 3 arguments: projectId, document, collection")
	}
}
