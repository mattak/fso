package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

func DeleteDoc(projectId, collection, doc string) *firestore.WriteResult {
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
	return result
}
