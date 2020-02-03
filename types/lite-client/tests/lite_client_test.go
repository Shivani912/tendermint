package tests

import (
	"fmt"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	generator "github.com/tendermint/tendermint/types/lite-client/generator"
)

func TestVerify(t *testing.T) {

	tests := []string{"commit_tests.json", "header_tests.json", "val_set_tests.json"}

	// tests := []string{"single_step_skipping/val_set_tests.json", "single_step_skipping/commit_tests.json"}
	for _, test := range tests {
		data := generator.ReadFile("./json/" + test)

		cdc := amino.NewCodec()
		cryptoAmino.RegisterAmino(cdc)

		var testBatch generator.TestBatch
		err := cdc.UnmarshalJSON(data, &testBatch)
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		for _, testCase := range testBatch.TestCases {

			chainID := testCase.Initial.SignedHeader.Header.ChainID
			trustedSignedHeader := testCase.Initial.SignedHeader
			trustedNextVals := testCase.Initial.NextValidatorSet
			trustingPeriod := testCase.Initial.TrustingPeriod
			now := testCase.Initial.Now
			trustLevel := lite.DefaultTrustLevel
			expectedOutput := testCase.ExpectedOutput
			expectsError := expectedOutput == "error"

			for _, input := range testCase.Input {

				fmt.Println(testCase.Description)
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
