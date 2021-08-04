package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"pitaya-serialization-benchmark/protos"

	"github.com/golang/protobuf/proto"
)

func main() {
	small := &protos.FetchProtoDataResponse{
		AString: "small",
	}

	smallData, _ := proto.Marshal(small)

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
	jsonFileData, err := os.ReadFile("large.json")
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

	large := &protos.FetchProtoDataResponse{
		Weapons: weapons,
	}

	largeData, err := proto.Marshal(large)
	if err != nil {
		panic(err)
	}

	return largeData
}
