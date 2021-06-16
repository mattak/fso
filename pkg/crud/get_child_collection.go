package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

type FirebaseCollectionMeta struct {
	ID          string `json:"id"`
	Path        string `json:"path"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
	ReadTime    string `json:"read_time"`
}

func CreateFirebaseCollectionMeta(snapshot *firestore.DocumentSnapshot) FirebaseCollectionMeta {
	return FirebaseCollectionMeta{
		ID:          snapshot.Ref.ID,
		Path:        snapshot.Ref.Path,
		CreatedTime: snapshot.CreateTime.String(),
		UpdatedTime: snapshot.UpdateTime.String(),
		ReadTime:    snapshot.ReadTime.String(),
	}
}

func GetChildCollectionMeta(projectId string, collection string, doc string) FirebaseCollectionMeta {
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

func GetChildCollections(projectId string, collection string, doc string) map[string]interface{} {
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
