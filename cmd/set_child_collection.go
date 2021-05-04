package cmd

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readFromStdin() interface{} {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func SetChildCollection(projectId, collection, doc string) {
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
	fmt.Println(result)
}
