package tests

import (
	"fmt"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	generator "github.com/tendermint/tendermint/types/lite-client/generator"
)

// TODO: deal with these globals
var cases generator.TestCases
var testCase generator.TestCase

func TestCase(t *testing.T) {

	data := generator.GetJsonFrom("./json/test_lite_client.json")

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)

	er := cdc.UnmarshalJSON(data, &cases)
	if er != nil {
		fmt.Printf("error: %v", er)
	}

	for _, tc := range cases.TC {

		testCase = tc

		switch testCase.Test {

		case "verify":
			t.Run("verify", TestVerify)
		default:
			fmt.Println("No such test found: ", testCase.Test)

		}
	}

}

func TestVerify(t *testing.T) {

	// TODO: clean up the arguments to these functions, eg:
	// signedHeader := testCase.Initial.SignedHeader
	// and then use signedHeader directly as necessary, etc.
	// Otherwise it's hard to read

	// TODO: deduplicate this logic by having some variable to refer to the latest trusted state.

	for i, input := range testCase.Input {
		if i == 0 {
			e := lite.Verify(testCase.Initial.SignedHeader.Header.ChainID, testCase.Initial.SignedHeader, &testCase.Initial.NextValidatorSet, input.SignedHeader, &input.ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)
			if e != nil {
				if e.Error() != testCase.ExpectedOutput[0] {
					t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testCase.Description, e, testCase.ExpectedOutput[0])
				}
			}
		} else {
			e := lite.Verify(testCase.Input[i-1].SignedHeader.Header.ChainID, testCase.Input[i-1].SignedHeader, &testCase.Input[i-1].NextValidatorSet, input.SignedHeader, &input.ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)
			if e != nil {
				if e.Error() != testCase.ExpectedOutput[i] {
					t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testCase.Description, e, testCase.ExpectedOutput[i])
				}
			}
		}
	}
}
