package main

import (
	"io/ioutil"
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

	large := &protos.FetchProtoDataResponse{
		AString: "large",
	}

	largeData, _ := proto.Marshal(large)

	ioutil.WriteFile("small.pb", smallData, 0644)
	ioutil.WriteFile("medium.pb", mediumData, 0644)
	ioutil.WriteFile("large.pb", largeData, 0644)
}
