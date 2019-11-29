package main

import "github.com/tendermint/tendermint/types/lite-client/generator"

func main() {

	generator.GenerateTestCases("./tests/json/val_list.json")
}
