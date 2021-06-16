package crud

import (
	"cloud.google.com/go/firestore"
	"context"
)

func ListUpRootCollections(projectId string) map[string]string {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	all, err := client.Collections(ctx).GetAll()
	if err != nil {
		panic(err)
	}

	result := map[string]string{}
	for i := 0; i < len(all); i++ {
		result[all[i].ID] = all[i].Path
	}
	return result
}
