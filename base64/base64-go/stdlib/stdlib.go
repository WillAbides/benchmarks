package main

import (
	"benchmarks/base64/base64-go"
	"encoding/base64"
)

func main() {
	err := base64_go.Run("", base64.StdEncoding)
	if err != nil {
		panic(err)
	}
}
