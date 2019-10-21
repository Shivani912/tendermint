package garbage

import (
	"fmt"
	"sort"
	"io/ioutil"

	"github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/lite2"


	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	st "github.com/tendermint/tendermint/state"

)

var state st.State
var testCase TestCase
var liteBlocks []*LiteBlock
var currentPrivVal types.PrivValidatorsByAddress

type TestCase struct {
	Name 			string 				`json:"name"`
	Initial			*LiteBlock			`json:"initial"`
	Input			[]*LiteBlock 		`json:"input"`
	ExpectedOutput	lite2.error			`json:"expected_output"`
}

type LiteBlock struct {
	SignedHeader *types.SignedHeader `json:"signed_header"`
	ValidatorSet types.ValidatorSet  `josn:"validator_set"`
}

func GenerateTestCase() {

	GenerateName("verify")
	GenerateFirstBlock()
	GenerateNextBlockWithNextValsUpdate(1, 10)
	GenerateNextBlock()
	GenerateExpectedOutput()
	GenerateJSON()
}

func GenerateStateFromGenesis() *st.State {

	state, err := st.MakeGenesisStateFromFile("./genesis.json")
	if err != nil {
		fmt.Println(err)
	}
	return &state
}

func GenerateFirstBlock() {
	newState := GenerateStateFromGenesis()
	state = newState.Copy()

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	valSet, privVal := types.RandValidatorSet(3, 7)
	state.Validators = valSet
	state.NextValidators = valSet

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, nil, evidences, state.Validators.Proposer.Address) // nil for last commit

	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, privVal)

	updateState(commit.BlockID, privVal)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *valSet,
	}

	testCase.Initial = lb
	liteBlocks = append(liteBlocks, lb)

}

func updateState(blockID types.BlockID, privVal []types.PrivValidator) {
	state.LastBlockHeight += 1
	state.LastValidators = state.Validators.Copy()
	state.Validators = state.NextValidators.Copy()
	state.Validators.IncrementProposerPriority(1)
	state.LastBlockID = blockID
	if len(currentPrivVal) != len(privVal) {
		for i:=0; i<len(privVal);i++ {
			currentPrivVal = append(currentPrivVal, privVal[i])
		}
		sort.Sort(currentPrivVal)
	}
	
}

func GenerateNextBlockWithNextValsUpdate(numVals int, votingPower int64) {

	valSet, privVal := types.RandValidatorSet(numVals, votingPower)
	err := state.NextValidators.UpdateWithChangeSet(valSet.Validators)
	if err != nil {
		fmt.Println(err)
	}
	state.NextValidators.IncrementProposerPriority(1)

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, liteBlocks[state.LastBlockHeight-1].SignedHeader.Commit, evidences, state.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, currentPrivVal)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *state.Validators,
	}
	liteBlocks = append(liteBlocks, lb)
	updateState(commit.BlockID, privVal)

}

func GenerateNextBlock() {

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()
	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, liteBlocks[state.LastBlockHeight-1].SignedHeader.Commit, evidences, state.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, currentPrivVal)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *state.Validators,
	}

	liteBlocks = append(liteBlocks, lb)
	updateState(commit.BlockID, currentPrivVal)

}

func GenerateJSON() {
	testCase.Input = liteBlocks[1:]

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)

	b, err := cdc.MarshalJSONIndent(testCase, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	file := "./test_case.json"
	_ = ioutil.WriteFile(file, b, 0644)

}

func GenerateName(name string) {
	testCase.Name = name
}

func GenerateExpectedOutput() {
	testCase.ExpectedOutput = nil
}
