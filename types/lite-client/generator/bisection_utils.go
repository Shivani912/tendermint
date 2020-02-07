package generator

import (
	"errors"
	"fmt"
	"time"

	lite "github.com/tendermint/tendermint/lite2"

	"github.com/tendermint/tendermint/types"
)

type TestBisection struct {
	Description    string            `json:"description"`
	TrustOptions   lite.TrustOptions `json:"trust_options"`
	Primary        MockProvider      `json:"primary"`
	HeightToVerify int64             `json:"height_to_verify"`
	TrustLevel     Fraction          `json:"trust_level"`
	Now            time.Time         `json:"now"`
	ExpectedOutput string            `json:"expected_output"`
}

type Fraction struct {
	// The portion of the denominator in the faction, e.g. 2 in 2/3.
	Numerator int64 `json:"numerator"`
	// The value by which the numerator is divided, e.g. 3 in 2/3. Must be
	// positive.
	Denominator int64 `json:"denominator"`
}

func (fr Fraction) String() string {
	return fmt.Sprintf("%d/%d", fr.Numerator, fr.Denominator)
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
	return nil, errors.New("sh not found")
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
	return nil, errors.New("vs not found")
}
