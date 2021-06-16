package crud

import (
	"encoding/json"
	"fmt"
)

func PrintJson(raw interface{}) {
	bytes, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	fmt.Println(data)
}

func PrintMapTsv(list map[string]string) {
	for k, v := range list {
		fmt.Printf("%s\t%s\n", k, v)
	}
}
