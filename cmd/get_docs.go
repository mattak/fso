package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
)

type rawMap map[string]interface{}

func GetDocs(projectId string, collection string) {
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
	resultMap := map[string]rawMap{}
	for i := 0; i < len(docs); i++ {
		id := docs[i].Ref.ID
		result := docs[i].Data()
		resultMap[id] = result
	}

	PrintJson(resultMap)
}
