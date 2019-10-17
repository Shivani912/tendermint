package main

import (
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"

	"github.com/tendermint/tendermint/version"
	"github.com/tendermint/tendermint/types"
	cmn "github.com/tendermint/tendermint/libs/common"

	"bufio"
    "fmt"
    "os"
)

// type Header struct {
// 	Height				int64			`json:"height"`
// 	Time				time.Time		`json:"time"`
// 	ValidatorsHash		cmn.HexBytes	`json:"validators_hash"`
// 	NextValidatorsHash	cmn.HexBytes	`json:"next_validators_hash"`
// }

// func generateLiteHeader() {

// 	var header Header
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Printf("Height: ")
// 	scanner.Scan()
// 	h := scanner.Text()
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading standard input:", err)

// 	height, err := strconv.ParseInt(h, 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	folder := "../"

// 	_, validatorsHash, _ := getData(folder)
	

// 	header = Header{
// 		Height: height,
// 		ValidatorsHash: validatorsHash,
// 		// NextValidatorsHash: nextValidatorsHash,
// 	}

// 	jsonHeader, err := json.MarshalIndent(header, "", "	")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
	
// 		file := "./header.json"
// 		_ = ioutil.WriteFile(file, jsonHeader, 0644)
// 		// fmt.Println(valHash)
// 	}
// }

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	q := [...]string{
		"blockVersion",
		"appVersion",
		"chainID",
		"totalTxs", 
		"valHash", 
		"nextValHash", 
		"consensusHash", 
		"appHash", 
		"lastResultsHash", 
		"proposerAddress",
	}

	var m map[string]string
	m = make(map[string]string)

	for i:=0;i<len(q);i++ {

		fmt.Printf(q[i]+": ")
		scanner.Scan()
		m[q[i]] = scanner.Text()
	}
	
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	

	var v version.Consensus

	// block, err := strconv.ParseUint(m["blockVersion"], 10, 64)
	// if err != nil {
    //     fmt.Println(err)
	// }
	// app, err := strconv.ParseUint(m["appVersion"], 10, 64)
	// if err != nil {
    //     fmt.Println(err)
    // }
	// v.Block = version.Protocol(block)
	// v.App = version.Protocol(app)

	// ChainID string
	chainID := m["chainID"]

	// Timestamp time.Time
	var timestamp time.Time
	timestamp = time.Now()

	// BlockID 
	var lastBlockID types.BlockID  
	lastBlockID = types.MakeBlockIDRandom()

	// TotalTxs int64
	t, err := strconv.ParseInt(m["totalTxs"], 10, 64)
	if err != nil {
        fmt.Println(err)
    }
	totalTxs := t


	// Validators Hash cmn.HexBytes
	valHash := []byte(m["valHash"])

	// Next validators Hash cmn.HexBytes
	nextValHash := cmn.HexBytes(m["nextValHash"])

	// Consensus Hash cmn.HexBytes
	consensusHash := cmn.HexBytes(m["consensusHash"])

	// App Hash cmn.HexBytes
	appHash := cmn.HexBytes(m["appHash"])

	// LastResultsHash cmn.HexBytes
	lastResultsHash := cmn.HexBytes(m["lastResultsHash"])

	// proposerAddress types.Address
	proposerAddress := types.Address(m["proposerAddress"])

	var h types.Header
	
	h.Populate(v, chainID, timestamp, lastBlockID, totalTxs, valHash, nextValHash, consensusHash, appHash, lastResultsHash, proposerAddress)


	header, err := json.MarshalIndent(h, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	file := "./header.json"
	_ = ioutil.WriteFile(file, header, 0644)
	// fmt.Println(valHash)
}
