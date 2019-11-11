package main

import (
	generators "github.com/tendermint/tendermint/types/lite-client/generators"
)

func main() {

	generators.GenerateTestCase("./tests/json/val_list.json")
}
