package garbage

import (
	"fmt"
	"io/ioutil"
	"sort"
	"time"

	"github.com/tendermint/tendermint/types"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	st "github.com/tendermint/tendermint/state"
)

var State st.State
var Case TestCase
var LiteBlocks []*LiteBlock
var CurrentPrivVal types.PrivValidatorsByAddress
var Cases TestCases

// var InitialData Initial

type TestCases struct {
	TC []TestCase `json:"test_cases"`
}

type TestCase struct {
	Test           string       `json:"name"`
	Description    string       `json:"description"`
	Initial        Initial      `json:"initial"`
	Input          []*LiteBlock `json:"input"`
	ExpectedOutput error        `json:"expected_output"`
}

type LiteBlock struct {
	SignedHeader *types.SignedHeader `json:"signed_header"`
	ValidatorSet types.ValidatorSet  `json:"validator_set"`
}

type Initial struct {
	SignedHeader     *types.SignedHeader `json:"signed_header"`
	NextValidatorSet types.ValidatorSet  `json:"next_validator_set"`
	TrustingPeriod   time.Duration       `json:"trusting_period"`
	Now              time.Time           `json:"now"`
}

func GenerateTestCase() {

	GenerateTestName("verify")
	Case.Description = "Case 1: Set of two lite blocks with no error"
	GenerateFirstBlock(1, 10)
	GenerateNextBlock()
	GenerateInitial(3*time.Hour, time.Now())
	Case.Input = LiteBlocks[1:]
	GenerateExpectedOutput()
	Cases.TC = append(Cases.TC, Case)

	ResetState()

	GenerateTestName("verify")
	Case.Description = "Case 2: Set of three lite blocks where validator set changes with no error."
	GenerateFirstBlock(1, 10)
	GenerateNextBlockWithNextValsUpdate(1, 10)
	GenerateInitial(3*time.Hour, time.Now())
	GenerateNextBlock()
	Case.Input = LiteBlocks[1:]
	GenerateExpectedOutput()
	Cases.TC = append(Cases.TC, Case)

	GenerateJSON()
}

func GenerateState() *st.State {

	consensusParams := types.ConsensusParams{
		Block:     types.DefaultBlockParams(),
		Evidence:  types.DefaultEvidenceParams(),
		Validator: types.DefaultValidatorParams(),
	}

	return &st.State{
		ChainID:         "test-chain-01",
		LastBlockHeight: 0,
		LastBlockID:     types.BlockID{},
		LastBlockTime:   time.Now(),

		NextValidators:              types.NewValidatorSet(nil),
		Validators:                  types.NewValidatorSet(nil),
		LastValidators:              types.NewValidatorSet(nil),
		LastHeightValidatorsChanged: 1,

		ConsensusParams:                  consensusParams,
		LastHeightConsensusParamsChanged: 1,

		AppHash: []byte("app_hash"),
	}
}

func ResetState() {
	LiteBlocks = nil
	State = st.State{}
	Case = TestCase{}
	CurrentPrivVal = nil
}

func GenerateFirstBlock(numVals int, votingPower int64) {

	newState := GenerateState()
	State = newState.Copy()

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	valSet, privVal := types.RandValidatorSet(numVals, votingPower)
	State.Validators = valSet
	State.NextValidators = valSet

	block, partSet := State.MakeBlock(State.LastBlockHeight+1, txs, nil, evidences, State.Validators.Proposer.Address) // nil for last commit

	commit := types.GenerateCommit(block.Header, partSet, *State.Validators, privVal, State.ChainID)

	updateState(commit.BlockID, privVal)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *valSet,
	}

	LiteBlocks = append(LiteBlocks, lb)

}

func updateState(blockID types.BlockID, privVal []types.PrivValidator) {
	State.LastBlockHeight += 1
	State.LastValidators = State.Validators.Copy()
	State.Validators = State.NextValidators.Copy()
	State.Validators.IncrementProposerPriority(1)
	State.LastBlockID = blockID

	for _, pv := range privVal {
		if !Contains(pv) {
			CurrentPrivVal = append(CurrentPrivVal, pv)
		}
	}
	sort.Sort(CurrentPrivVal)
}

func Contains(pv types.PrivValidator) bool {
	for _, n := range CurrentPrivVal {
		if pv == n {
			return true
		}
	}
	return false
}

func GenerateInitial(trustingPeriod time.Duration, now time.Time) {

	Case.Initial.SignedHeader = LiteBlocks[0].SignedHeader
	Case.Initial.NextValidatorSet = LiteBlocks[1].ValidatorSet
	Case.Initial.TrustingPeriod = trustingPeriod
	Case.Initial.Now = now
}

func GenerateNextBlockWithNextValsUpdate(numVals int, votingPower int64) {

	valSet, privVal := types.RandValidatorSet(numVals, votingPower)
	err := State.NextValidators.UpdateWithChangeSet(valSet.Validators)
	if err != nil {
		fmt.Println(err)
	}
	State.NextValidators.IncrementProposerPriority(1)

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	block, partSet := State.MakeBlock(State.LastBlockHeight+1, txs, LiteBlocks[State.LastBlockHeight-1].SignedHeader.Commit, evidences, State.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *State.Validators, CurrentPrivVal, State.ChainID)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *State.Validators,
	}
	LiteBlocks = append(LiteBlocks, lb)
	updateState(commit.BlockID, privVal)

}

func GenerateNextBlock() {

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()
	block, partSet := State.MakeBlock(State.LastBlockHeight+1, txs, LiteBlocks[State.LastBlockHeight-1].SignedHeader.Commit, evidences, State.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *State.Validators, CurrentPrivVal, State.ChainID)

	lb := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet: *State.Validators,
	}

	LiteBlocks = append(LiteBlocks, lb)
	updateState(commit.BlockID, CurrentPrivVal)

}

func GenerateJSON() {

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)
	cdc.RegisterInterface((*error)(nil), nil)

	b, err := cdc.MarshalJSONIndent(Cases, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	file := "./test_lite_client.json"
	_ = ioutil.WriteFile(file, b, 0644)

}

func GenerateTestName(testName string) {
	Case.Test = testName
}

func GenerateExpectedOutput() {
	e := lite.Verify(Case.Initial.SignedHeader.Header.ChainID, Case.Initial.SignedHeader, &Case.Initial.NextValidatorSet, Case.Input[0].SignedHeader, &Case.Input[0].ValidatorSet, Case.Initial.TrustingPeriod, Case.Initial.Now, lite.DefaultTrustLevel)

	Case.ExpectedOutput = e
}
