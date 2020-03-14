package generator

import (
	"fmt"
	"time"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmmath "github.com/tendermint/tendermint/libs/math"
	st "github.com/tendermint/tendermint/state"

	"github.com/tendermint/tendermint/lite2/provider"
	"github.com/tendermint/tendermint/types"
)

type TestBisection struct {
	Description    string              `json:"description"`
	TrustOptions   TrustOptions        `json:"trust_options"`
	Primary        MockProvider        `json:"primary"`
	Witnesses      []provider.Provider `json:"witnesses"`
	HeightToVerify int64               `json:"height_to_verify"`
	Now            time.Time           `json:"now"`
	ExpectedOutput string              `json:"expected_output"`
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
	for _, lb := range mp.LiteBlocks {
		if lb.SignedHeader.Header.Height == height {
			return &lb.ValidatorSet, nil
		}
		if lb.SignedHeader.Header.Height+1 == height {
			return &lb.NextValidatorSet, nil
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
	for i := 0; i < numOfBlocks; i++ {
		liteblock, st, pvs := generateNextBlockWithNextValsUpdate(
			state,
			privVals,
			lastCommit,
			valSetChanges[i].Validators,
			valSetChanges[i].PrivVals,
			blockTime,
		)
		liteBlocks = append(liteBlocks, liteblock)
		state = st
		privVals = pvs
		lastCommit = liteblock.SignedHeader.Commit
		states = append(states, state)
		blockTime = blockTime.Add(5 * time.Second)
	}
	return liteBlocks, states, privVals
}

type ValSetChanges []ValList

func (vsc ValSetChanges) makeValSetChangeAtHeight(
	height int,
	vals []*types.Validator,
	privVals types.PrivValidatorsByAddress,
) {
	vsc[height] = ValList{
		Validators: vals,
		PrivVals:   privVals,
	}
}

func (vsc *ValSetChanges) makeValSetChanges(
	vals [][]*types.Validator,
	privVals []types.PrivValidatorsByAddress,
) {
	for i := range vals {
		v := ValList{
			Validators: vals[i],
			PrivVals:   privVals[i],
		}
		*vsc = append(*vsc, v)
	}
}
