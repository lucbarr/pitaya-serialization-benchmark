package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	jsonData, err := ioutil.ReadFile(filename)
	fmt.Println(filename)
	if err != nil {
		panic(err)
	}

	jsonMap := map[string]interface{}{}

	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		panic(err)
	}

	halfNumberOfKeys := len(jsonMap) / 2

	halvedJsonMap := map[string]interface{}{}
	i := 0
	for key, value := range jsonMap {
		halvedJsonMap[key] = value
		i++
		if i > halfNumberOfKeys {
			break
		}
	}

	halvedJsonData, err := json.Marshal(halvedJsonMap)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("halved-%s", filename), halvedJsonData, 0644)
	if err != nil {
		panic(err)
	}
}
