package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func ListUpDocs(projectId string, collection string) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	statues := client.Collection(collection)
	docs, err := statues.Documents(ctx).GetAll()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(docs); i++ {
		fmt.Printf("%s\t%s\n", docs[i].Ref.ID, docs[i].Ref.Path)
	}
}
