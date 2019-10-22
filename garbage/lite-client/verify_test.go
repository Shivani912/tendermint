package main 

import (
	"fmt"
	"time"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	generator "github.com/tendermint/tendermint/garbage"
	lite "github.com/tendermint/tendermint/lite2"
)

func TestVerify(t *testing.T) {
	data, err := getJsonFrom("./test_case.json")
	if err != 0 {
		fmt.Println(err)
	}

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*error)(nil), nil)

	var testCase generator.TestCase
	er := cdc.UnmarshalJSON(data, &testCase)
	if er != nil {
		fmt.Printf("error: %v", er)
	}

	e := lite.Verify(testCase.Initial.SignedHeader.Header.ChainID, testCase.Initial.SignedHeader, &testCase.Input[0].ValidatorSet, testCase.Input[0].SignedHeader, &testCase.Input[0].ValidatorSet, 3 * time.Hour, time.Now(), lite.DefaultTrustLevel)
	
	if e != testCase.ExpectedOutput {
		t.Errorf("\n Error: %v", e)
	}
}