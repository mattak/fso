package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func CreateChildCollection(projectId, collection, doc string) {
	data := readFromStdin()

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	statues := client.Collection(collection)
	collectionRef := statues.Doc(doc)

	result, err := collectionRef.Create(ctx, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
