package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func DeleteChildCollection(projectId, collection, doc string) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	statues := client.Collection(collection)
	collectionRef := statues.Doc(doc)

	result, err := collectionRef.Delete(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
