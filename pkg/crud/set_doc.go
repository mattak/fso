package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

func SetDoc(projectId, collection, doc string) *firestore.WriteResult {
	data := readFromStdin()

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	statues := client.Collection(collection)
	collectionRef := statues.Doc(doc)

	result, err := collectionRef.Set(ctx, data, firestore.MergeAll)
	if err != nil {
		panic(err)
	}
	return result
}
