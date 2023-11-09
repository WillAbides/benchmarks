package main

import (
	"benchmarks/common/json/json-go"

	"github.com/willabides/rjson"
)

func calc(bytes []byte) (json_go.Coordinate, error) {
	obj := json_go.TestStruct{}
	var buffer rjson.Buffer
	var currentCoord *json_go.Coordinate

	coordHandler := rjson.ObjectValueHandlerFunc(func(fieldname, data []byte) (int, error) {
		switch string(fieldname) {
		case "x":
			return rjson.DecodeFloat64(data, &currentCoord.X)
		case "y":
			return rjson.DecodeFloat64(data, &currentCoord.Y)
		case "z":
			return rjson.DecodeFloat64(data, &currentCoord.Z)
		}
		return rjson.SkipValueFast(data, &buffer)
	})

	coordsHandler := rjson.ArrayValueHandlerFunc(func(data []byte) (int, error) {
		obj.Coordinates = append(obj.Coordinates, json_go.Coordinate{})
		currentCoord = &obj.Coordinates[len(obj.Coordinates)-1]
		return rjson.HandleObjectValues(data, coordHandler, &buffer)
	})

	_, err := rjson.HandleObjectValues(bytes, rjson.ObjectValueHandlerFunc(func(fieldname, data []byte) (int, error) {
		if string(fieldname) != "coordinates" {
			return rjson.SkipValueFast(data, &buffer)
		}
		return rjson.HandleArrayValues(data, coordsHandler, &buffer)
	}), &buffer)
	if err != nil {
		return json_go.Coordinate{}, err
	}
	return obj.Average(), nil
}

func main() {
	err := json_go.Run("rjson", calc)
	if err != nil {
		panic(err)
	}
}
