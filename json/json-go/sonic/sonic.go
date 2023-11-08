package main

import (
	"benchmarks/common/json/json-go"

	"github.com/bytedance/sonic"
)

func calc(bytes []byte) (json_go.Coordinate, error) {
	obj := json_go.TestStruct{}
	err := sonic.Unmarshal(bytes, &obj)
	if err != nil {
		return json_go.Coordinate{}, err
	}
	return obj.Average(), nil
}

func main() {
	err := json_go.Run("Sonic", calc)
	if err != nil {
		panic(err)
	}
}
