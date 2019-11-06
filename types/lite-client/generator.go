package main

import (
	generators "github.com/tendermint/tendermint/types/lite-client/generators"
)

func main() {

	generators.GenerateTestCase("./val_list.json")
}
