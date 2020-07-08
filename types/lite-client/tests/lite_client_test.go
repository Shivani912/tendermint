package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	"github.com/tendermint/tendermint/lite2/provider"
	"github.com/tendermint/tendermint/types/lite-client/generator"

	dbs "github.com/tendermint/tendermint/lite2/store/db"
	dbm "github.com/tendermint/tm-db"
)

func TestVerify(t *testing.T) {

	tests, err := getTestPaths("./json/single_step/")
	if err != nil {
		fmt.Println(err)
	}

	for _, test := range tests {
		data := generator.ReadFile(test)

		cdc := amino.NewCodec()
		cryptoAmino.RegisterAmino(cdc)

		var testCase generator.TestCase
		err := cdc.UnmarshalJSON(data, &testCase)
		if err != nil {
			fmt.Printf("error: %v", err)
		}

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
			fmt.Printf("\n%s, \nError: %v \n", testCase.Description, e)
			if (err && !expectsError) || (!err && expectsError) {
				t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testCase.Description, e, expectedOutput)

			} else {
				trustedSignedHeader = newSignedHeader
				trustedNextVals = input.NextValidatorSet
			}
		}
	}

}

func TestBisection(t *testing.T) {
	tests, err := getTestPaths("./json/bisection/")
	if err != nil {
		fmt.Println(err)
	}

	for _, test := range tests {
		data := generator.ReadFile(test)

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
		// ---
		fmt.Println(e)
		// ---
		err := e != nil
		expectsError := expectedOutput == "error"
		if (err && !expectsError) || (!err && expectsError) {
			t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testBisection.Description, e, testBisection.ExpectedOutput)

		}
	}
}

func getTestPaths(folder string) ([]string, error) {
	var tests []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			tests = append(tests, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}
	return tests, nil
}
