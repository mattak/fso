package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

func ListUpDocs(projectId string, collection string) map[string]string {
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

	result := map[string]string{}
	for i := 0; i < len(docs); i++ {
		result[docs[i].Ref.ID] = docs[i].Ref.Path
	}
	return result
}
