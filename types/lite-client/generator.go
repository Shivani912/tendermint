package main

import "github.com/tendermint/tendermint/types/lite-client/generator"

func main() {

	generator.GenerateTestCase("./tests/json/val_list.json")
}
