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

	signedHeader, state, privVals := generateFirstBlock(
		valList.Validators[0:3],
		valList.PrivVals[0:3],
		firstBlockTime,
	)

	trustOptions := TrustOptions{
		Period:     trustingPeriod,
		Height:     signedHeader.Header.Height,
		Hash:       signedHeader.Commit.BlockID.Hash,
		TrustLevel: lite.DefaultTrustLevel,
	}

	firstBlock := LiteBlock{
		SignedHeader:     signedHeader,
		ValidatorSet:     *state.Validators,
		NextValidatorSet: *state.NextValidators,
	}

	lastCommit := signedHeader.Commit

	valsArray := [][]*types.Validator{
		valList.Validators[:4],
		valList.Validators[:5],
		valList.Validators[:6],
		valList.Validators[:7],
		valList.Validators[:8],
		valList.Validators[5:9],
		valList.Validators[5:10],
		valList.Validators[5:11],
		valList.Validators[5:12],
		valList.Validators[5:13],
	}
	privValsArray := []types.PrivValidatorsByAddress{
		valList.PrivVals[:4],
		valList.PrivVals[:5],
		valList.PrivVals[:6],
		valList.PrivVals[:7],
		valList.PrivVals[:8],
		valList.PrivVals[5:9],
		valList.PrivVals[5:10],
		valList.PrivVals[5:11],
		valList.PrivVals[5:12],
		valList.PrivVals[5:13],
	}

	var valSetChanges ValSetChanges
	valSetChanges.makeValSetChanges(valsArray, privValsArray)
	lbs, _, _ := generateNextBlocks(10, state, privVals, lastCommit, valSetChanges, thirdBlockTime)

	primary := MockProvider{}.New(signedHeader.Header.ChainID, []LiteBlock{})

	primary.LiteBlocks = append(primary.LiteBlocks, firstBlock)
	primary.LiteBlocks = append(primary.LiteBlocks, lbs...)

	var witnesses []provider.Provider
	witnesses = append([]provider.Provider{}, primary)

	heightToVerify := int64(11)

	expectedOutput := expectedOutputNoError

	testBisection := TestBisection{
		Description:    description,
		TrustOptions:   trustOptions,
		Primary:        primary,
		Witnesses:      witnesses,
		HeightToVerify: heightToVerify,
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
