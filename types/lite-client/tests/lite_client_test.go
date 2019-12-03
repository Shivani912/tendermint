package tests

import (
	"fmt"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	generator "github.com/tendermint/tendermint/types/lite-client/generator"
)

// DONE: deal with these globals

func TestVerify(t *testing.T) {

	// DONE: clean up the arguments to these functions, eg:
	// signedHeader := testCase.Initial.SignedHeader
	// and then use signedHeader directly as necessary, etc.
	// Otherwise it's hard to read

	// DONE: deduplicate this logic by having some variable to refer to the latest trusted state.

	tests := []string{"commit_tests.json", "header_tests.json", "val_set_tests.json"}
	for _, test := range tests {
		data := generator.ReadFile("./json/" + test)

		cdc := amino.NewCodec()
		cryptoAmino.RegisterAmino(cdc)

		var testCases generator.TestCases
		err := cdc.UnmarshalJSON(data, &testCases)
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		for _, testCase := range testCases.TC {

			chainID := testCase.Initial.SignedHeader.Header.ChainID
			trustedSignedHeader := testCase.Initial.SignedHeader
			trustedNextVals := testCase.Initial.NextValidatorSet
			trustingPeriod := testCase.Initial.TrustingPeriod
			now := testCase.Initial.Now
			trustLevel := lite.DefaultTrustLevel
			expectedOutput := testCase.ExpectedOutput
			expectsError := expectedOutput == "error"

			for _, input := range testCase.Input {

				newSignedHeader := input.SignedHeader
				newVals := input.ValidatorSet

				e := lite.Verify(chainID, &trustedSignedHeader, &trustedNextVals, &newSignedHeader, &newVals, trustingPeriod, now, trustLevel)
				err := e != nil

				if (err && !expectsError) || (!err && expectsError) {
					t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testCase.Description, e, expectedOutput)

				} else {
					trustedSignedHeader = newSignedHeader
					trustedNextVals = input.NextValidatorSet
				}
			}
		}
	}

}
