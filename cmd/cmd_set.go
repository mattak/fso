package cmd

import (
	"fmt"
	"github.com/mattak/fso/pkg/crud"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	SetCmd = &cobra.Command{
		Use:     "set [PROJECT_ID] [COLLECTION_ID] [DOC_ID]",
		Aliases: []string{"s"},
		Short:   "Set data",
		Long:    "Set data",
		Example: `  echo '{"symbols":["a"]}' | crud set projectId collection1 document1
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

	result := crud.SetDoc(projectId, collection, doc)
	fmt.Println(result)
}
