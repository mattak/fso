package cmd

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
