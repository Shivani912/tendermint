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
	description := "Case: Trusted height=1, bisecting to verify height=11, should not expect error"

	signedHeader, state, privVals := generateFirstBlock(valList, 3, firstBlockTime)

	trustOptions := lite.TrustOptions{
		Period: trustingPeriod,
		Height: signedHeader.Header.Height,
		Hash:   signedHeader.Header.Hash(),
	}

	firstBlock := LiteBlock{
		SignedHeader:     signedHeader,
		ValidatorSet:     *state.Validators,
		NextValidatorSet: *state.NextValidators,
	}

	lastCommit := signedHeader.Commit

	// TODO: Improve the way we make changes to valset
	// possibly a `ValSetChanges` struct that includes the below info
	start := []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	end := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	delete := []int{0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0}
	lbs, _, _ := generateNextBlocks(10, state, privVals, lastCommit, valList, start, end, delete, thirdBlockTime)

	primary := MockProvider{}.New(signedHeader.Header.ChainID, []LiteBlock{})
	primary.LiteBlocks = append(primary.LiteBlocks, firstBlock)
	primary.LiteBlocks = append(primary.LiteBlocks, lbs...)

	heightToVerify := int64(11)
	trustLevel := Fraction{
		Numerator:   1,
		Denominator: 3,
	}

	expectedOutput := expectedOutputNoError

	testBisection := TestBisection{
		Description:    description,
		TrustOptions:   trustOptions,
		Primary:        primary,
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
