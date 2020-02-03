package main

import "github.com/tendermint/tendermint/types/lite-client/generator"

func main() {

	generator.GenerateSingleStepSequentialCases("./tests/json/val_list.json")
	// tlcOutput := generator.GetTLCOutput("./generator/tlc_output.txt")
	// valList := generator.GetValList("./tests/json/val_list.json")
	// generator.TlcOutputToTestCase(tlcOutput, valList)
}
