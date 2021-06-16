package crud

import (
	"encoding/json"
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
