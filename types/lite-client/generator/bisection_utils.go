package generator

import (
	"fmt"
	"io/ioutil"
	"time"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmmath "github.com/tendermint/tendermint/libs/math"
	lite "github.com/tendermint/tendermint/lite2"
	st "github.com/tendermint/tendermint/state"

	"github.com/tendermint/tendermint/lite2/provider"
	"github.com/tendermint/tendermint/types"
)

type TestBisection struct {
	Description        string              `json:"description"`
	TrustOptions       TrustOptions        `json:"trust_options"`
	Primary            MockProvider        `json:"primary"`
	Witnesses          []provider.Provider `json:"witnesses"`
	HeightToVerify     int64               `json:"height_to_verify"`
	Now                time.Time           `json:"now"`
	ExpectedOutput     string              `json:"expected_output"`
	ExpectedBisections int32               `json:"expected_num_of_bisections"`
}

func (tb TestBisection) make(
	desc string,
	trustOpts TrustOptions,
	primary MockProvider,
	witnesses []provider.Provider,
	heightToVerify int64,
	now time.Time,
	expectedOutput string,
	expectedBisections int32,
) TestBisection {
	return TestBisection{
		Description:        desc,
		TrustOptions:       trustOpts,
		Primary:            primary,
		Witnesses:          witnesses,
		HeightToVerify:     heightToVerify,
		Now:                now,
		ExpectedOutput:     expectedOutput,
		ExpectedBisections: expectedBisections,
	}
}

func (testBisection TestBisection) genJSON(file string) {
	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)
	cdc.RegisterInterface((*provider.Provider)(nil), nil)
	cdc.RegisterConcrete(MockProvider{}, "com.tendermint/MockProvider", nil)

	b, err := cdc.MarshalJSONIndent(testBisection, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	_ = ioutil.WriteFile(file, b, 0644)
}

type TrustOptions struct {
	// Trusting Period
	Period time.Duration `json:"period"`
	// Trusted Header Height
	Height int64 `json:"height"`
	// Trusted Header Hash
	Hash       tmbytes.HexBytes `json:"hash"`
	TrustLevel tmmath.Fraction  `json:"trust_level"`
}

func (t TrustOptions) make(
	sh types.SignedHeader,
	trustingPeriod time.Duration,
	trustLevel tmmath.Fraction,
) TrustOptions {
	return TrustOptions{
		Period:     trustingPeriod,
		Height:     sh.Header.Height,
		Hash:       sh.Commit.BlockID.Hash,
		TrustLevel: trustLevel,
	}

}

type MockProvider struct {
	ChainId    string      `json:"chain_id"`
	LiteBlocks []LiteBlock `json:"lite_blocks"`
}

func (mp MockProvider) New(chainID string, liteBlocks []LiteBlock) MockProvider {
	return MockProvider{
		ChainId:    chainID,
		LiteBlocks: liteBlocks,
	}
}

func (mp MockProvider) Copy() MockProvider {
	return MockProvider{
		ChainId:    mp.ChainId,
		LiteBlocks: mp.LiteBlocks,
	}
}

func (mp MockProvider) ChainID() string {
	return mp.ChainId
}

func (mp MockProvider) SignedHeader(height int64) (*types.SignedHeader, error) {
	fmt.Printf("\n sh -- req h: %v", height)
	for _, lb := range mp.LiteBlocks {
		if lb.SignedHeader.Header.Height == height {
			return &lb.SignedHeader, nil
		}
	}
	return nil, provider.ErrSignedHeaderNotFound
}
func (mp MockProvider) ValidatorSet(height int64) (*types.ValidatorSet, error) {
	fmt.Printf("\n vs -- req h: %v", height)
	// if lb.SignedHeader.Header.Height+1 == height {
	// 		return &lb.NextValidatorSet, nil
	// 	}
	for _, lb := range mp.LiteBlocks {
		if lb.SignedHeader.Header.Height == height {
			return &lb.ValidatorSet, nil
		}
	}
	return nil, provider.ErrValidatorSetNotFound
}

func generateNextBlocks(
	numOfBlocks int,
	state st.State,
	privVals types.PrivValidatorsByAddress,
	lastCommit *types.Commit,
	valSetChanges ValSetChanges,
	blockTime time.Time,
) ([]LiteBlock, []st.State, types.PrivValidatorsByAddress) {
	var liteBlocks []LiteBlock
	var states []st.State
	valSetChanges = append(valSetChanges, valSetChanges[len(valSetChanges)-1])
	for i := 0; i < numOfBlocks; i++ {
		liteblock, st, _ := generateNextBlockWithNextValsUpdate(
			state,
			valSetChanges[i].PrivVals,
			lastCommit,
			valSetChanges[i+1].Validators,
			nil,
			blockTime,
		)
		liteBlocks = append(liteBlocks, liteblock)
		state = st
		lastCommit = liteblock.SignedHeader.Commit
		states = append(states, state)
		blockTime = blockTime.Add(5 * time.Second)
	}
	return liteBlocks, states, privVals
}

type ValSetChanges []ValList

func (vsc ValSetChanges) getDefault(valList ValList) ValSetChanges {
	valsArray := [][]*types.Validator{
		valList.Validators[:2],
		valList.Validators[:2],
		valList.Validators[:2],
		valList.Validators[:2],
		valList.Validators[:2],
		valList.Validators[3:5],
		valList.Validators[3:5],
		valList.Validators[3:5],
		valList.Validators[3:5],
		valList.Validators[3:5],
		valList.Validators[3:5],
	}
	privValsArray := []types.PrivValidatorsByAddress{
		valList.PrivVals[:2],
		valList.PrivVals[:2],
		valList.PrivVals[:2],
		valList.PrivVals[:2],
		valList.PrivVals[:2],
		valList.PrivVals[3:5],
		valList.PrivVals[3:5],
		valList.PrivVals[3:5],
		valList.PrivVals[3:5],
		valList.PrivVals[3:5],
		valList.PrivVals[3:5],
	}
	return vsc.makeValSetChanges(valsArray, privValsArray)
}

func (vsc ValSetChanges) makeValSetChangeAtHeight(
	height int,
	vals []*types.Validator,
	privVals types.PrivValidatorsByAddress,
) ValSetChanges {
	vsc[height] = ValList{
		Validators: vals,
		PrivVals:   privVals,
	}
	return vsc
}

func (vsc ValSetChanges) makeValSetChanges(
	vals [][]*types.Validator,
	privVals []types.PrivValidatorsByAddress,
) ValSetChanges {
	for i := range vals {
		v := ValList{
			Validators: vals[i],
			PrivVals:   privVals[i],
		}
		vsc = append(vsc, v)
	}
	return vsc
}

func makeLiteblocks(
	valSetChanges ValSetChanges,
) ([]LiteBlock, []st.State, types.PrivValidatorsByAddress) {
	signedHeader, state, _ := generateFirstBlockWithNextValsUpdate(
		valSetChanges[0].Validators,
		valSetChanges[0].PrivVals,
		valSetChanges[1].Validators,
		nil,
		firstBlockTime,
	)

	firstBlock := []LiteBlock{
		{
			SignedHeader:     signedHeader,
			ValidatorSet:     *state.LastValidators,
			NextValidatorSet: *state.Validators,
		},
	}
	lastCommit := signedHeader.Commit
	numOfBlocks := len(valSetChanges) - 1
	liteBlocks, states, privVals := generateNextBlocks(
		numOfBlocks,
		state,
		valSetChanges[1].PrivVals,
		lastCommit,
		valSetChanges[1:],
		thirdBlockTime,
	)
	liteBlocks = append(firstBlock, liteBlocks...)
	stateSlice := []st.State{
		state,
	}
	states = append(stateSlice, states...)
	return liteBlocks, states, privVals
}

func generateMultiPeerBisectionCase(
	description string,
	primaryValSetChanges ValSetChanges,
	alternativeValSetChanges ValSetChanges,
	expectedBisections int32,
	expectOutput string,
) (TestBisection, []st.State, types.PrivValidatorsByAddress, []st.State, types.PrivValidatorsByAddress) {
	testBisection, statesPrimary, privValsPrimary := generateGeneralBisectionCase(
		description,
		primaryValSetChanges,
		expectedBisections)

	liteBlocks, statesAlternative, privValsAlternative := makeLiteblocks(alternativeValSetChanges)
	testBisection.Witnesses[0] = MockProvider{}.New(liteBlocks[0].SignedHeader.Header.ChainID, liteBlocks)
	testBisection.ExpectedOutput = expectOutput
	return testBisection, statesPrimary, privValsPrimary, statesAlternative, privValsAlternative
}

func generateGeneralBisectionCase(
	description string,
	valSetChanges ValSetChanges,
	expectedBisections int32,
) (TestBisection, []st.State, types.PrivValidatorsByAddress) {

	liteBlocks, states, privVals := makeLiteblocks(valSetChanges)
	primary := MockProvider{}.New(liteBlocks[0].SignedHeader.Header.ChainID, liteBlocks)

	var witnesses []provider.Provider
	witnesses = append([]provider.Provider{}, primary)

	trustOptions := TrustOptions{}.make(liteBlocks[0].SignedHeader, TRUSTING_PERIOD, lite.DefaultTrustLevel)
	heightToVerify := int64(len(valSetChanges))

	testBisection := TestBisection{}.make(
		description,
		trustOptions,
		primary,
		witnesses,
		heightToVerify,
		now,
		expectedOutputNoError,
		expectedBisections,
	)

	return testBisection, states, privVals
}
