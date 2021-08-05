package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pitaya-serialization-benchmark/protos"

	"github.com/golang/protobuf/proto"

	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func main() {
	smallData := generateSmallProtoData()

	medium := &protos.FetchProtoDataResponse{
		AString: "medium",
	}

	mediumData, _ := proto.Marshal(medium)

	largeData := generateLargeProtoData()

	ioutil.WriteFile("small.pb", smallData, 0644)
	ioutil.WriteFile("medium.pb", mediumData, 0644)
	ioutil.WriteFile("large.pb", largeData, 0644)
}

func generateLargeProtoData() []byte {
	jsonFileData, err := os.ReadFile("../json/large.json")
	if err != nil {
		panic(err)
	}

	weapons := &protos.Weapons{
		Weapons: map[string]*protos.Weapons_Weapon{},
	}

	err = json.Unmarshal(jsonFileData, &weapons.Weapons)
	if err != nil {
		panic(err)
	}

	validateGeneratedProto(jsonFileData, weapons.Weapons)

	large := &protos.FetchProtoDataResponse{
		Weapons: weapons,
	}

	largeData, err := proto.Marshal(large)
	if err != nil {
		panic(err)
	}

	return largeData
}

func generateSmallProtoData() []byte {
	jsonFileData, err := os.ReadFile("../json/small.json")
	if err != nil {
		panic(err)
	}

	gears := &protos.Gears{}

	err = json.Unmarshal(jsonFileData, gears)
	if err != nil {
		panic(err)
	}

	small := &protos.FetchProtoDataResponse{
		Gears: gears,
	}

	smallData, err := proto.Marshal(small)
	if err != nil {
		panic(err)
	}

	return smallData
}

var formatterConfig = formatter.AsciiFormatterConfig{
	ShowArrayIndex: true,
	Coloring:       true,
}

func validateGeneratedProto(sourceJSONData []byte, proto interface{}) {
	retJSONFromProto, err := json.Marshal(proto)
	if err != nil {
		panic(err)
	}

	differ := diff.New()

	diff, err := differ.Compare(sourceJSONData, retJSONFromProto)
	if err != nil {
		panic(err)
	}

	if diff.Modified() {
		expectedJSON := map[string]interface{}{}
		json.Unmarshal(sourceJSONData, &expectedJSON)
		formatter := formatter.NewAsciiFormatter(expectedJSON, formatterConfig)
		diffString, err := formatter.Format(diff)
		if err != nil {
			panic(err)
		}
		fmt.Println(diffString)
	}
}
