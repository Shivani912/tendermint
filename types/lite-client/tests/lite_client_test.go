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

	tests := []string{
		"single_step_sequential/commit_tests.json",
		"single_step_sequential/header_tests.json",
		"single_step_sequential/val_set_tests.json",
		"single_step_skipping/val_set_tests.json",
		"single_step_skipping/commit_tests.json",
	}

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

				newSignedHeader := input.SignedHeader
				newVals := input.ValidatorSet

				e, str := lite.Verify(chainID, &trustedSignedHeader, &trustedNextVals, &newSignedHeader, &newVals, trustingPeriod, now, trustLevel)
				err := e != nil
				fmt.Printf("\n%s, \nError: %v \n", testCase.Description, e)
				if (err && !expectsError) || (!err && expectsError) {
					t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v \n String: %v", testCase.Description, e, expectedOutput, str)

				} else {
					trustedSignedHeader = newSignedHeader
					trustedNextVals = input.NextValidatorSet
				}
			}
		}
	}

}

func TestBisection(t *testing.T) {
	tests := []string{
		"many_header_bisection/happy_path.json",
		"many_header_bisection/worst_case.json",
		"many_header_bisection/invalid_validator_set.json",
		"many_header_bisection/not_enough_commits.json",
		"many_header_bisection/header_out_of_trusting_period.json",
	}

	for _, test := range tests {
		data := generator.ReadFile("./json/" + test)

		cdc := amino.NewCodec()
		cryptoAmino.RegisterAmino(cdc)

		cdc.RegisterInterface((*provider.Provider)(nil), nil)
		cdc.RegisterConcrete(generator.MockProvider{}, "com.tendermint/MockProvider", nil)

		var testBisection generator.TestBisection
		e := cdc.UnmarshalJSON(data, &testBisection)
		if e != nil {
			fmt.Printf("error: %v", e)
		}

		fmt.Println(testBisection.Description)

		trustedStore := dbs.New(dbm.NewMemDB(), testBisection.Primary.ChainID())
		witnesses := testBisection.Witnesses
		trustOptions := lite.TrustOptions{
			Period: testBisection.TrustOptions.Period,
			Height: testBisection.TrustOptions.Height,
			Hash:   testBisection.TrustOptions.Hash,
		}
		trustLevel := testBisection.TrustOptions.TrustLevel
		expectedOutput := testBisection.ExpectedOutput

		client, e := lite.NewClient(
			testBisection.Primary.ChainID(),
			trustOptions,
			testBisection.Primary,
			witnesses,
			trustedStore,
			lite.SkippingVerification(trustLevel))
		if e != nil {
			fmt.Println(e)
		}

		height := testBisection.HeightToVerify
		_, e = client.VerifyHeaderAtHeight(height, testBisection.Now)

		err := e != nil
		expectsError := expectedOutput == "error"
		if (err && !expectsError) || (!err && expectsError) {
			t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testBisection.Description, e, testBisection.ExpectedOutput)

		}
	}
}
