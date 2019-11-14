package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"

	lite "github.com/tendermint/tendermint/lite2"
	"github.com/tendermint/tendermint/types"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	st "github.com/tendermint/tendermint/state"
)

var genTime, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:00Z")
var now, _ = time.Parse(time.RFC3339, "2019-11-02T15:05:00Z")
var firstBlockTime, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:10Z")
var secondBlockTime, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:15Z")
var thirdBlockTime, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:20Z")

type TestCases struct {
	TC []TestCase `json:"test_cases"`
}

type TestCase struct {
	Test           string       `json:"name"`
	Description    string       `json:"description"`
	Initial        Initial      `json:"initial"`
	Input          []*LiteBlock `json:"input"`
	ExpectedOutput string       `json:"expected_output"`
}

type LiteBlock struct {
	SignedHeader     *types.SignedHeader `json:"signed_header"`
	ValidatorSet     types.ValidatorSet  `json:"validator_set"`
	NextValidatorSet types.ValidatorSet  `json:"next_validator_set"`
}

type Initial struct {
	SignedHeader     *types.SignedHeader `json:"signed_header"`
	NextValidatorSet types.ValidatorSet  `json:"next_validator_set"`
	TrustingPeriod   time.Duration       `json:"trusting_period"`
	Now              time.Time           `json:"now"`
}

type ValList struct {
	ValidatorSet types.ValidatorSet            `json:"validator_set"`
	PrivVal      types.PrivValidatorsByAddress `json:"priv_val"`
}

func NewState() *st.State {

	consensusParams := types.ConsensusParams{
		Block:     types.DefaultBlockParams(),
		Evidence:  types.DefaultEvidenceParams(),
		Validator: types.DefaultValidatorParams(),
	}

	return &st.State{
		ChainID:         "test-chain-01",
		LastBlockHeight: 0,
		LastBlockID:     types.BlockID{},
		LastBlockTime:   genTime,

		NextValidators:              types.NewValidatorSet(nil),
		Validators:                  types.NewValidatorSet(nil),
		LastValidators:              types.NewValidatorSet(nil),
		LastHeightValidatorsChanged: 1,

		ConsensusParams:                  consensusParams,
		LastHeightConsensusParamsChanged: 1,

		AppHash: []byte("app_hash"),
	}
}

func GenerateFirstBlock(testCase *TestCase, valz []*types.Validator, privVal types.PrivValidatorsByAddress, now time.Time) *st.State {

	state := NewState()

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	valSet := types.NewValidatorSet(valz)
	state.Validators = valSet
	state.NextValidators = valSet

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, nil, evidences, state.Validators.Proposer.Address) // nil for last commit

	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, privVal, state.ChainID, now)

	uState, _ := updateState(state, commit.BlockID, privVal, nil)

	initial := Initial{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
	}
	testCase.Initial = initial

	return uState
}

func updateState(state *st.State, blockID types.BlockID, privVal types.PrivValidatorsByAddress, newPrivVal types.PrivValidatorsByAddress) (*st.State, types.PrivValidatorsByAddress) {
	state.LastBlockHeight += 1
	state.LastValidators = state.Validators.Copy()
	state.Validators = state.NextValidators.Copy()
	state.Validators.IncrementProposerPriority(1)
	state.LastBlockID = blockID

	for _, npv := range newPrivVal {
		if !Contains(privVal, npv) {
			privVal = append(privVal, npv)
		}
	}

	if len(privVal) > len(state.Validators.Validators) {
		for i := 0; i < len(privVal); i++ {
			_, val := state.Validators.GetByAddress(privVal[i].GetPubKey().Address())
			if val == nil {
				privVal = append(privVal[:i], privVal[i+1:]...)
				i = i - 1
			}

		}
	}

	sort.Sort(privVal)

	return state, privVal
}

func Contains(privVal types.PrivValidatorsByAddress, npv types.PrivValidator) bool {
	for _, n := range privVal {
		if npv == n {
			return true
		}
	}
	return false
}

func GenerateInitial(testCase *TestCase, nextValidatorSet types.ValidatorSet, trustingPeriod time.Duration, now time.Time) {

	testCase.Initial.NextValidatorSet = nextValidatorSet
	testCase.Initial.TrustingPeriod = trustingPeriod
	testCase.Initial.Now = now

}

func GenerateNextBlockWithNextValsUpdate(testCase *TestCase, state *st.State, privVal types.PrivValidatorsByAddress, lastCommit *types.Commit, newVals []*types.Validator, newPrivVal types.PrivValidatorsByAddress, delete int, now time.Time) types.PrivValidatorsByAddress {

	if delete > 0 {
		for i := 0; i < delete; i++ {
			state.NextValidators.Validators[i].VotingPower = 0
			newVals = append(newVals, state.NextValidators.Validators[i])
		}
	}
	err := state.NextValidators.UpdateWithChangeSet(newVals)
	if err != nil {
		fmt.Println(err)
	}
	state.NextValidators.IncrementProposerPriority(1)

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, lastCommit, evidences, state.Validators.Proposer.Address)
	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, privVal, state.ChainID, now)

	liteBlock := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet:     *state.Validators.Copy(),
		NextValidatorSet: *state.NextValidators.Copy(),
	}
	testCase.Input = append(testCase.Input, liteBlock)
	state, uPrivVal := updateState(state, commit.BlockID, privVal, newPrivVal)

	return uPrivVal
}

func GenerateNextBlock(state *st.State, testCase *TestCase, privVal types.PrivValidatorsByAddress, lastCommit *types.Commit, now time.Time) {

	txs := types.GenerateTxs()
	evidences := types.GenerateEvidences()

	block, partSet := state.MakeBlock(state.LastBlockHeight+1, txs, lastCommit, evidences, state.Validators.Proposer.Address)

	commit := types.GenerateCommit(block.Header, partSet, *state.Validators, privVal, state.ChainID, now)
	liteBlock := &LiteBlock{
		SignedHeader: &types.SignedHeader{
			Header: &block.Header,
			Commit: commit,
		},
		ValidatorSet:     *state.Validators.Copy(),
		NextValidatorSet: *state.NextValidators.Copy(),
	}

	testCase.Input = append(testCase.Input, liteBlock)

	uState, _ := updateState(state, commit.BlockID, privVal, privVal)
	state = uState

}

func GenerateJSON(testCases *TestCases) {

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)

	b, err := cdc.MarshalJSONIndent(testCases, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	file := "./tests/json/test_lite_client.json"
	_ = ioutil.WriteFile(file, b, 0644)

}

func GenerateTestNameAndDescription(testCase *TestCase, testName string, description string) {

	testCase.Test = testName
	testCase.Description = description
}

func GenerateExpectedOutput(testCase *TestCase) {
	for i, input := range testCase.Input {
		if i == 0 {
			e := lite.Verify(testCase.Initial.SignedHeader.Header.ChainID, testCase.Initial.SignedHeader, &testCase.Initial.NextValidatorSet, input.SignedHeader, &input.ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)
			if e != nil {
				testCase.ExpectedOutput = e.Error()
			}
		} else {
			e := lite.Verify(testCase.Input[i-1].SignedHeader.Header.ChainID, testCase.Input[i-1].SignedHeader, &testCase.Input[i-1].NextValidatorSet, input.SignedHeader, &input.ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)
			if e != nil {
				testCase.ExpectedOutput = e.Error()
			}
		}

	}

}

func GenerateValList(numVals int, votingPower int64) {

	valSet, newPrivVal := types.RandValidatorSet(numVals, votingPower)
	sort.Sort(types.ValidatorsByAddress(valSet.Validators))
	valList := &ValList{
		ValidatorSet: *valSet,
		PrivVal:      types.PrivValidatorsByAddress(newPrivVal),
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

func GetJsonFrom(file string) []byte {
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

func GetValList(file string) ValList {
	data := GetJsonFrom(file)
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
