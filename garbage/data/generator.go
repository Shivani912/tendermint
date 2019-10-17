package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"

	// cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/types"

)

func main() {
	txs := types.MakeTxs(5, 10)

	jsonTxs, err := json.MarshalIndent(txs, " ", "	")
	if err != nil {
		fmt.Println(err)
	}

	file := "./data.json"
	_ = ioutil.WriteFile(file, jsonTxs, 0644)

}