package main

import (
	// "fmt"
	"time"
	"strconv"
	"crypto/rand"

	"github.com/tendermint/tendermint/version"
	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	cmn "github.com/tendermint/tendermint/libs/common"

	"bufio"
    "fmt"
    "os"
)

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
	

	// Assign values
	
	// Version struct {
	// 	Block version.Protocol, 
	// 	App version.Protocol
	// }
	var v version.Consensus

	block, err := strconv.ParseUint(m["blockVersion"], 10, 64)
	if err != nil {
        fmt.Println(err)
	}
	app, err := strconv.ParseUint(m["appVersion"], 10, 64)
	if err != nil {
        fmt.Println(err)
    }
	v.Block = version.Protocol(block)
	v.App = version.Protocol(app)

	// ChainID string
	chainID := m["chainID"]

	// Timestamp time.Time
	var timestamp time.Time
	timestamp = time.Now()

	// BlockID 
	var lastBlockID types.BlockID  
	lastBlockID = makeBlockIDRandom()

	// TotalTxs int64
	t, err := strconv.ParseInt(m["totalTxs"], 10, 64)
	if err != nil {
        fmt.Println(err)
    }
	totalTxs := t

	// Validators Hash cmn.HexBytes
	valHash := cmn.HexBytes(m["valHash"])

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

	fmt.Println(valHash)
}

func makeBlockIDRandom() types.BlockID {
	blockHash := make([]byte, tmhash.Size)
	partSetHash := make([]byte, tmhash.Size)
	rand.Read(blockHash)   //nolint: gosec
	rand.Read(partSetHash) //nolint: gosec
	blockPartsHeader := types.PartSetHeader{123, partSetHash}
	return types.BlockID{blockHash, blockPartsHeader}
}