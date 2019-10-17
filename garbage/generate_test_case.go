package garbage

import (
	"fmt"
	"io/ioutil"

	"github.com/tendermint/tendermint/types"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	st "github.com/tendermint/tendermint/state"
)

var state st.State

var blockSlice []*types.Block
var liteBlocks []*types.LiteBlock
var currentPrivVal []types.PrivValidator

func GenerateTestCase() {

	GenerateFirstBlock()

	GenerateNextBlockWithNextValsUpdate(4, 7)

	GenerateNextBlock()

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

	lb := &types.LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *valSet,
	}

	liteBlocks = append(liteBlocks, lb)
	// blockSlice = append(blockSlice, block)

}

func updateState(blockID types.BlockID, privVal []types.PrivValidator) {
	state.LastBlockHeight += 1
	state.LastValidators = state.Validators.Copy()
	state.Validators = state.NextValidators.Copy()
	state.LastBlockID = blockID
	currentPrivVal = privVal
}

func GenerateNextBlockWithNextValsUpdate(numVals int, votingPower int64) {

	valSet, privVal := types.RandValidatorSet(numVals, votingPower)
	err := state.NextValidators.UpdateWithChangeSet(valSet.Validators)
	if err != nil {
		fmt.Println(err)
	}

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, liteBlocks[state.LastBlockHeight-1].SignedHeader.Commit, evidences, state.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, currentPrivVal)

	lb := &types.LiteBlock{
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
	fmt.Printf("\nVALIDATORS---- \n %v \nPRIVVAL \n %v", state.Validators, currentPrivVal)
	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, liteBlocks[state.LastBlockHeight-1].SignedHeader.Commit, evidences, state.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, currentPrivVal)

	lb := &types.LiteBlock{
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
	testCase := &types.TestCase{
		LiteBlocks: liteBlocks,
	}

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
