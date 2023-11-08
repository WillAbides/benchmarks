package main

import (
	"benchmarks/base64/base64-go"
	"github.com/chenzhuoyu/base64x"
)

func main() {
	err := base64_go.Run("base64x", base64x.StdEncoding)
	if err != nil {
		panic(err)
	}
}
