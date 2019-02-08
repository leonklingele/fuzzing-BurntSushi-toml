package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type empty struct{}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var e empty
	_, _ = toml.Decode(string(data), &e)
}
