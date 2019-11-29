package main

import "github.com/tendermint/tendermint/types/lite-client/generator"

func main() {

	generator.GenerateTestCases("./tests/json/val_list.json", "./tests/json/test_lite_client.json")
}
