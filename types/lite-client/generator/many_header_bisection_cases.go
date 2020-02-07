package generator

import (
	"fmt"
	"io/ioutil"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	"github.com/tendermint/tendermint/lite2/provider"
	"github.com/tendermint/tendermint/types"
)

func caseBisectionVerifyTenHeaders(valList ValList) {
	description := "Case: Trusted height=1, bisecting to verify height=12, should not expect error"

	signedHeader, state, privVals := generateFirstBlock(valList, 3, firstBlockTime)

	trustOptions := lite.TrustOptions{
		Period: trustingPeriod,
		Height: signedHeader.Header.Height,
		Hash:   signedHeader.Commit.BlockID.Hash,
	}

	firstBlock := LiteBlock{
		SignedHeader:     signedHeader,
		ValidatorSet:     *state.Validators,
		NextValidatorSet: *state.NextValidators,
	}

	lastCommit := signedHeader.Commit
	lbs, _, _ := generateNextBlocks(10, state, privVals, lastCommit)

	primary := MockProvider{}.New(signedHeader.Header.ChainID, []LiteBlock{})

	primary.LiteBlocks = append(primary.LiteBlocks, firstBlock)
	primary.LiteBlocks = append(primary.LiteBlocks, lbs...)

	var witnesses []provider.Provider
	witnesses = append([]provider.Provider{}, primary)

	heightToVerify := int64(11)
	trustLevel := lite.DefaultTrustLevel

	expectedOutput := expectedOutputNoError

	testBisection := TestBisection{
		Description:    description,
		TrustOptions:   trustOptions,
		Primary:        primary,
		Witnesses:      witnesses,
		HeightToVerify: heightToVerify,
		TrustLevel:     trustLevel,
		Now:            now,
		ExpectedOutput: expectedOutput,
	}

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*types.Evidence)(nil), nil)
	cdc.RegisterInterface((*provider.Provider)(nil), nil)
	cdc.RegisterConcrete(MockProvider{}, "com.tendermint/MockProvider", nil)

	b, err := cdc.MarshalJSONIndent(testBisection, " ", "	")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	_ = ioutil.WriteFile("./tests/json/many_header_bisection/happy_path.json", b, 0644)

}
