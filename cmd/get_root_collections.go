package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
)

func GetRootCollections(projectId string) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}

	all, err := client.Collections(ctx).GetAll()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(all); i++ {
		fmt.Printf("%s\t%s\n", all[i].ID, all[i].Path)
	}
}
