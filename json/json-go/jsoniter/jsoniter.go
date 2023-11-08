package main

import (
	"benchmarks/common/json/json-go"

	jsoniter "github.com/json-iterator/go"
)

func calc(bytes []byte) (json_go.Coordinate, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	obj := json_go.TestStruct{}
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return json_go.Coordinate{}, err
	}
	return obj.Average(), nil
}

func main() {
	err := json_go.Run("jsoniter", calc)
	if err != nil {
		panic(err)
	}
}
