package main

import (
	"benchmarks/common/json/json-go"

	"encoding/json"
)

func calc(bytes []byte) (json_go.Coordinate, error) {
	obj := json_go.TestStruct{}
	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return json_go.Coordinate{}, err
	}
	return obj.Average(), nil
}

func main() {
	err := json_go.Run("", calc)
	if err != nil {
		panic(err)
	}
}
