package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/tendermint/tendermint/types"
)

func main() {
	voteSet := types.ExamplePrevote()

	privVal := types.NewMockPV()
	err := privVal.SignVote("test_chain_id", voteSet)

	jsonVote, err := json.MarshalIndent(voteSet, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	file := "./vote.json"
	_ = ioutil.WriteFile(file, jsonVote, 0644)
	fmt.Println("Vote.json created!")
}