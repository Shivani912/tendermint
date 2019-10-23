package main

import (
	"fmt"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	generator "github.com/tendermint/tendermint/garbage"
	lite "github.com/tendermint/tendermint/lite2"
)

var cases generator.TestCases
var testCase generator.TestCase

func TestCase(t *testing.T) {

	data, err := getJsonFrom("./test_lite_client.json")
	if err != 0 {
		fmt.Println(err)
	}

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*error)(nil), nil)

	er := cdc.UnmarshalJSON(data, &cases)
	if er != nil {
		fmt.Printf("error: %v", er)
	}

	for _, tc := range cases.TC {

		testCase = tc

		switch testCase.Test {

		case "verify":
			t.Run("verify", TestVerify)
			// case "negative":
			// 	t.Run("Negative Case", NegativeCase)
		default:
			fmt.Println("No such test found: ", testCase.Test)

		}
	}

}

func TestVerify(t *testing.T) {

	e := lite.Verify(testCase.Initial.SignedHeader.Header.ChainID, testCase.Initial.SignedHeader, &testCase.Initial.NextValidatorSet, testCase.Input[0].SignedHeader, &testCase.Input[0].ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)

	if e != testCase.ExpectedOutput {
		t.Errorf("\n Failing test: %s \n Error: %v", testCase.Description, e)
	}
}
