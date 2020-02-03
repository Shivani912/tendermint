package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	"github.com/tendermint/tendermint/types"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	st "github.com/tendermint/tendermint/state"
)

var (
	genTime, _         = time.Parse(time.RFC3339, "2019-11-02T15:04:00Z")
	now, _             = time.Parse(time.RFC3339, "2019-11-02T15:30:00Z")
	firstBlockTime, _  = time.Parse(time.RFC3339, "2019-11-02T15:04:10Z")
	secondBlockTime, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:15Z")
	thirdBlockTime, _  = time.Parse(time.RFC3339, "2019-11-02T15:04:20Z")
	trustingPeriod     = 3 * time.Hour
	// testName              = "verify"
	expectedOutputNoError = "no error"
	expectedOutputError   = "error"
)

// TestBatch contains a slice of TestCase, for now.
// It may contain other information in future
type TestBatch struct {
	BatchName string     `json:"batch_name"`
	TestCases []TestCase `json:"test_cases"`
}

// TestCase stores all the necessary information to perform the test on the data given
type TestCase struct {
	Description    string      `json:"description"`
	Initial        Initial     `json:"initial"`
	Input          []LiteBlock `json:"input"`
	ExpectedOutput string      `json:"expected_output"`
}

// LiteBlock refers to the minimum data a lite client interacts with.
// Essentially, it only requires a SignedHeader and Validator Set for current and next height
type LiteBlock struct {
	SignedHeader     types.SignedHeader `json:"signed_header"`
	ValidatorSet     types.ValidatorSet `json:"validator_set"`
	NextValidatorSet types.ValidatorSet `json:"next_validator_set"`
}

// Initial stores the data required by a test case to set the context
// i.e. the initial state to begin the test from
type Initial struct {
	SignedHeader     types.SignedHeader `json:"signed_header"`
	NextValidatorSet types.ValidatorSet `json:"next_validator_set"`
	TrustingPeriod   time.Duration      `json:"trusting_period"`
	Now              time.Time          `json:"now"`
}

// ValList stores a list of validators and privVals
// It is populated from the lite-client/tests/json/val_list.json
// It used to have a predefined set of validators for mocking the test data
type ValList struct {
	Validators []*types.Validator            `json:"validators"`
	PrivVals   types.PrivValidatorsByAddress `json:"priv_val"`
}

// NewState is used to initiate a state that will be used and manipulated
// by functions to create blocks for the "simulated" blockchain
// It creates an INITIAL state with the given parameters
func NewState(chainID string, valSet *types.ValidatorSet) st.State {

	consensusParams := types.ConsensusParams{
		Block:     types.DefaultBlockParams(),
		Evidence:  types.DefaultEvidenceParams(),
		Validator: types.DefaultValidatorParams(),
	}

	return st.State{
		ChainID:         chainID,
		LastBlockHeight: 0,
		LastBlockID:     types.BlockID{},
		LastBlockTime:   genTime,

		NextValidators:              valSet,
		Validators:                  valSet,
		LastValidators:              types.NewValidatorSet(nil),
		LastHeightValidatorsChanged: 1,

		ConsensusParams:                  consensusParams,
		LastHeightConsensusParamsChanged: 1,

		AppHash: []byte("app_hash"),
	}
}

// generateFirstBlock creates the first block of the chain
// with the given list of validators and timestamp
// Thus, It also calls the NewState() to initialize the state
// Returns the signedHeader and state after the first block is created
func generateFirstBlock(valList ValList, numOfValz int, now time.Time) (types.SignedHeader, st.State, types.PrivValidatorsByAddress) {

	valz := valList.Validators[:numOfValz]
	privVals := valList.PrivVals[:numOfValz]

	valSet := types.NewValidatorSet(valz)
	state := NewState("test-chain-01", valSet)

	txs := generateTxs()
	evidences := generateEvidences()
	lbh := state.LastBlockHeight + 1
	proposer := state.Validators.Proposer.Address

	// first block has a nil last commit
	block, partSet := state.MakeBlock(lbh, txs, nil, evidences, proposer)

	commit := generateCommit(block.Header, partSet, state.Validators, privVals, state.ChainID, now)

	state, privVals = updateState(state, commit.BlockID, privVals, nil)

	signedHeader := types.SignedHeader{
		Header: &block.Header,
		Commit: commit,
	}

	return signedHeader, state, privVals
}

// Called after creating each block to update the validator set, proposer,
// last block id, privVals etc.
// In case of privVals, it adds the new ones to the list
// and performs a sort operation on it.
func updateState(state st.State, blockID types.BlockID, privVals types.PrivValidatorsByAddress, newPrivVals types.PrivValidatorsByAddress) (st.State, types.PrivValidatorsByAddress) {
	state.LastBlockHeight++
	state.LastValidators = state.Validators.Copy()
	state.Validators = state.NextValidators.Copy()
	state.Validators.IncrementProposerPriority(1)
	state.LastBlockID = blockID

	// Adds newPrivVals if they are not already present in privVals
	if newPrivVals != nil {
		for _, npv := range newPrivVals {
			if !contains(privVals, npv) {
				privVals = append(privVals, npv)
			}
		}
	}

	// Checks if a validator has been removed from the set
	// If so, removes it from privVals too
	if len(privVals) > len(state.Validators.Validators) {
		for i := 0; i < len(privVals); i++ {
			_, val := state.Validators.GetByAddress(privVals[i].GetPubKey().Address())
			if val == nil {
				// removing the privVal when no corresponding entry found in the validator set
				privVals = append(privVals[:i], privVals[i+1:]...)
				i = i - 1
			}
		}
	}

	sort.Sort(privVals)

	return state, privVals
}

// Checks if privVals contain the privVal - used by updateState()
func contains(privVals types.PrivValidatorsByAddress, npv types.PrivValidator) bool {
	for _, n := range privVals {
		if npv == n {
			return true
		}
	}
	return false
}

// Builds the Initial struct with given parameters
func generateInitial(signedHeader types.SignedHeader, nextValidatorSet types.ValidatorSet, trustingPeriod time.Duration, now time.Time) Initial {

	return Initial{
		SignedHeader:     signedHeader,
		NextValidatorSet: nextValidatorSet,
		TrustingPeriod:   trustingPeriod,
		Now:              now,
	}
}

// This one generates a "next" block,
// i.e. given the first block, this function can be used to build up successive blocks
func generateNextBlock(state st.State, privVals types.PrivValidatorsByAddress, lastCommit *types.Commit, now time.Time) (LiteBlock, st.State) {

	txs := generateTxs()
	evidences := generateEvidences()
	lbh := state.LastBlockHeight + 1
	proposer := state.Validators.Proposer.Address

	block, partSet := state.MakeBlock(lbh, txs, lastCommit, evidences, proposer)

	commit := generateCommit(block.Header, partSet, state.Validators, privVals, state.ChainID, now)
	liteBlock := LiteBlock{
		SignedHeader: types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet:     *state.Validators.Copy(),     // dereferencing pointer
		NextValidatorSet: *state.NextValidators.Copy(), // dereferencing pointer
	}

	state, _ = updateState(state, commit.BlockID, privVals, nil)
	return liteBlock, state

}

// Similar to generateNextBlock
// It also takes in new validators and privVals to be added to the NextValidatorSet
// Calls the UpdateWithChangeSet function on state.NextValidatorSet for the same
// Also, you can specify the number of vals to be deleted from it
func generateNextBlockWithNextValsUpdate(state st.State, privVals types.PrivValidatorsByAddress, lastCommit *types.Commit, valList ValList, startIdx int, endIdx int, delete int, now time.Time) (LiteBlock, st.State, types.PrivValidatorsByAddress) {

	copyValList := valList.Copy()
	newVals := copyValList.Validators[startIdx:endIdx]
	newPrivVals := copyValList.PrivVals[startIdx:endIdx]
	if delete > 0 && delete < len(state.NextValidators.Validators)+len(newVals) {
		for i := 0; i < delete; i++ {
			toDelete := *state.NextValidators.Validators[i]
			toDelete.VotingPower = 0
			newVals = append(newVals, &toDelete)
		}
	}
	err := state.NextValidators.UpdateWithChangeSet(newVals)
	if err != nil {
		fmt.Println(err)
	}
	state.NextValidators.IncrementProposerPriority(1)

	txs := generateTxs()
	evidences := generateEvidences()
	lbh := state.LastBlockHeight + 1
	proposer := state.Validators.Proposer.Address

	block, partSet := state.MakeBlock(lbh, txs, lastCommit, evidences, proposer)
	commit := generateCommit(block.Header, partSet, state.Validators, privVals, state.ChainID, now)

	liteBlock := LiteBlock{
		SignedHeader: types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet:     *state.Validators.Copy(),     // dereferencing pointer
		NextValidatorSet: *state.NextValidators.Copy(), // dereferencing pointer
	}
	state, newPrivVals = updateState(state, commit.BlockID, privVals, newPrivVals)

	return liteBlock, state, newPrivVals
}

// generateJSON produces the JSON for the given testCase type.
// The ouput is saved under the specified file parameter
func generateJSON(testCases *TestBatch, file string) {

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)

	b, err := cdc.MarshalJSONIndent(testCases, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	_ = ioutil.WriteFile(file, b, 0644)

}

// makeTestCase copies over the given parameters to the TestCase struct and returns it
func makeTestCase(description string, initial Initial, input []LiteBlock, expectedOutput string) TestCase {
	return TestCase{
		Description:    description,
		Initial:        initial,
		Input:          input,
		ExpectedOutput: expectedOutput,
	}
}

// GenerateValList produces a val_list.json file which contains a list validators and privVals
// of given number abd voting power
func GenerateValList(numVals int, votingPower int64) {

	valSet, newPrivVal := types.RandValidatorSet(numVals, votingPower)
	sort.Sort(types.ValidatorsByAddress(valSet.Validators))
	valList := &ValList{
		Validators: valSet.Validators,
		PrivVals:   types.PrivValidatorsByAddress(newPrivVal),
	}

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.PrivValidator)(nil), nil)
	cdc.RegisterConcrete(&types.MockPV{}, "tendermint/MockPV", nil)

	b, err := cdc.MarshalJSONIndent(valList, " ", "	")

	if err != nil {
		panic(err)
	}

	file := "./val_list.json"
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		panic(err)
	}
}

// ReadFile returns the byte slice of the content in the given file
// "file" parameter is the path to the file to be read
func ReadFile(file string) []byte {
	jsonFile, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	dat, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err)
	}

	return dat
}

// GetValList reads the validators and privals list in the file
// unmarshals it to ValList struct
// "file" parameter specifies the path to the val_list.json file
func GetValList(file string) ValList {
	data := ReadFile(file)
	var valList ValList

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.PrivValidator)(nil), nil)
	cdc.RegisterConcrete(&types.MockPV{}, "tendermint/MockPV", nil)

	er := cdc.UnmarshalJSON(data, &valList)
	if er != nil {
		fmt.Printf("error: %v", er)
	}

	return valList
}

// Builds a general case containing initial and one lite block in input
// TODO: change name to genInitialAndInput
func generateGeneralCase(valList ValList, numOfVals int) (Initial, []LiteBlock, st.State, types.PrivValidatorsByAddress) {

	var input []LiteBlock

	signedHeader, state, privVals := generateFirstBlock(valList, numOfVals, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)
	liteBlock, state := generateNextBlock(state, privVals, signedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)

	return initial, input, state, privVals
}

func generateInitialAndInputSkipBlocks(valList ValList, numOfVals, numOfBlocksToSkip int) (Initial, []LiteBlock, st.State, types.PrivValidatorsByAddress) {
	var input []LiteBlock

	signedHeader, state, privVals := generateFirstBlock(valList, numOfVals, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	blockTime := secondBlockTime
	for i := 0; i <= numOfBlocksToSkip; i++ {
		liteBlock, s := generateNextBlock(state, privVals, signedHeader.Commit, blockTime)
		blockTime = blockTime.Add(5 * time.Second)
		state = s

		if i == numOfBlocksToSkip {
			input = append(input, liteBlock)
		}
	}

	return initial, input, state, privVals
}

func generateAndMakeGeneralTestCase(description string, valList ValList, numOfVals int, expectedOutput string) TestCase {

	initial, input, _, _ := generateGeneralCase(valList, numOfVals)
	return makeTestCase(description, initial, input, expectedOutput)
}

func generateAndMakeNextValsUpdateTestCase(description string, valList ValList, numOfInitialVals int, numOfValsToAdd int, numOfValsToDelete int, expectedOutput string) TestCase {

	copyValList := valList.Copy()
	initial, input, _, _ := generateNextValsUpdateCase(copyValList, numOfInitialVals, numOfValsToAdd, numOfValsToDelete)
	return makeTestCase(description, initial, input, expectedOutput)
}

// Builds a case where next validator set changes
// makes initial, and input with 2 lite blocks
func generateNextValsUpdateCase(valList ValList, numOfInitialVals int, numOfValsToAdd int, numOfValsToDelete int) (Initial, []LiteBlock, st.State, types.PrivValidatorsByAddress) {

	var input []LiteBlock

	signedHeader, state, privVals := generateFirstBlock(valList, numOfInitialVals, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	startIdx := numOfInitialVals
	endIdx := startIdx + numOfValsToAdd
	liteBlock, state, privVals := generateNextBlockWithNextValsUpdate(state, privVals, signedHeader.Commit, valList, startIdx, endIdx, numOfValsToDelete, secondBlockTime)
	input = append(input, liteBlock)
	liteBlock, state = generateNextBlock(state, privVals, liteBlock.SignedHeader.Commit, thirdBlockTime)
	input = append(input, liteBlock)

	return initial, input, state, privVals
}

// UPDATE -> mutex on PartSet and functions take pointer to valSet - have to use a pointer
// generateCommit takes in header, partSet from Block that was created,
// validator set, privVals, chain ID and a timestamp to create
// and return a commit type
// To be called after MakeBlock()
func generateCommit(header types.Header, partSet *types.PartSet, valSet *types.ValidatorSet, privVals []types.PrivValidator, chainID string, now time.Time) *types.Commit {
	blockID := types.BlockID{
		Hash: header.Hash(),
		PartsHeader: types.PartSetHeader{
			Hash:  partSet.Hash(),
			Total: partSet.Total(),
		},
	}
	voteSet := types.NewVoteSet(chainID, header.Height, 1, types.SignedMsgType(byte(types.PrecommitType)), valSet)

	commit, err := types.MakeCommit(blockID, header.Height, 1, voteSet, privVals, now)
	if err != nil {
		fmt.Println(err)
	}

	return commit
}

func generateTxs() []types.Tx {
	// Empty txs
	return []types.Tx{}
}

func generateEvidences() []types.Evidence {
	// Empty evidences
	return []types.Evidence{}
}

// Copy is essentially used to dereference the pointer
// ValList contains valSet pointer and privVal interface
// So to avoid manipulating the original list, we better copy it!
func (valList ValList) Copy() (vl ValList) {

	for i, val := range valList.Validators {
		// var privVal types.PrivValidator
		copyVal := *val
		privVal := valList.PrivVals[i]
		vl.Validators = append(vl.Validators, &copyVal)
		vl.PrivVals = append(vl.PrivVals, privVal)
	}
	return
}
