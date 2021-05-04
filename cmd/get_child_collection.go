package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
)

func GetChildCollections(projectId string, collection string, doc string, isMeta bool) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	statues := client.Collection(collection)
	collectionRef := statues.Doc(doc)
	result, err := collectionRef.Get(ctx)
	if err != nil {
		panic(err)
	}

	if isMeta {
		maps := map[string]string{
			"ID":          result.Ref.ID,
			"Path":        result.Ref.Path,
			"CreatedTime": result.CreateTime.String(),
			"UpdateTime":  result.UpdateTime.String(),
			"ReadTime":    result.ReadTime.String(),
		}
		PrintJson(maps)
	} else {
		PrintJson(result.Data())
	}
}
