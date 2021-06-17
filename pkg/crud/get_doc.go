package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

type rawMap map[string]interface{}

func GetDocs(projectId string, collection string) map[string]rawMap {
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

	return resultMap
}

func GetDocMeta(projectId string, collection string, doc string) FirebaseCollectionMeta {
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
	return CreateFirebaseCollectionMeta(result)
}

func GetDoc(projectId string, collection string, doc string) map[string]interface{} {
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

	return result.Data()
}
