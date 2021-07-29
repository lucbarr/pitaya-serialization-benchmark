package services

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/topfreegames/pitaya/component"
	"golang.org/x/xerrors"
)

type BenchmarkHandler struct {
	component.Base
	dataExamplesFetcher DataExamplesFetcher
}

func NewBenchmarkHandler() *BenchmarkHandler {
	dataExamplesFetcher := NewDataExamplesRepository()

	return &BenchmarkHandler{
		dataExamplesFetcher: dataExamplesFetcher,
	}
}

type Size int
type DataType int

const (
	Small Size = iota
	Medium
	Large
)

var allSizes = []Size{
	Small,
	Medium,
	Large,
}

func (s Size) String() string {
	switch s {
	case Small:
		return "small"
	case Medium:
		return "medium"
	case Large:
		return "large"
	}

	return ""
}

const (
	JSON DataType = iota
	Protobuf
)

type DataExamplesFetcher interface {
	Fetch(size Size, dataType DataType) ([]byte, error)
}

type dataExamplesRepository struct {
	protos map[Size][]byte
	json   map[Size][]byte
}

func loadProtos() map[Size][]byte {
	loadedProtos := map[Size][]byte{}
	for _, size := range allSizes {
		filename := fmt.Sprintf("examples/proto/%s.pb", size.String())
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		loadedProtos[size] = data
	}

	return loadedProtos
}

func loadJSON() map[Size][]byte {
	loadedProtos := map[Size][]byte{}
	for _, size := range allSizes {
		filename := fmt.Sprintf("examples/json/%s.json", size.String())
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}

		loadedProtos[size] = data
	}

	return loadedProtos
}

func NewDataExamplesRepository() *dataExamplesRepository {
	loadedProtos := loadProtos()
	loadedJSON := loadJSON()

	return &dataExamplesRepository{
		protos: loadedProtos,
		json:   loadedJSON,
	}
}

var (
	ErrInvalidDataType error = xerrors.New("invalid data type")
	ErrInvalidSize     error = xerrors.New("invalid size")
)

func (d *dataExamplesRepository) Fetch(size Size, dataType DataType) ([]byte, error) {
	var data []byte
	var ok bool

	switch dataType {
	case Protobuf:
		data, ok = d.protos[size]
		break
	case JSON:
		data, ok = d.json[size]
		break
	default:
		return nil, ErrInvalidDataType
	}

	if !ok {
		return nil, ErrInvalidSize
	}

	return data, nil

}
