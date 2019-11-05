package main

import (
	"fmt"
	"testing"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	lite "github.com/tendermint/tendermint/lite2"
	generator "github.com/tendermint/tendermint/tests"
)

var cases generator.TestCases
var testCase generator.TestCase

func TestCase(t *testing.T) {

	data := generator.GetJsonFrom("./test_lite_client.json")

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	// cdc.RegisterInterface((*error)(nil), nil)

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

	e := lite.Verify(testCase.Initial.SignedHeader.Header.ChainID, testCase.Initial.SignedHeader, &testCase.Initial.NextValidatorSet, testCase.Input[0].SignedHeader, &testCase.Input[0].ValidatorSet, testCase.Initial.TrustingPeriod, testCase.Initial.Now, lite.DefaultTrustLevel)

	if e != nil {
		if e.Error() != testCase.ExpectedOutput {
			t.Errorf("\n Failing test: %s \n Error: %v \n Expected error: %v", testCase.Description, e, testCase.ExpectedOutput)
		}
	}
}

// func TestValList(t *testing.T) {
// 	data := getJsonFrom("./val_list.json")

// 	var cdc = amino.NewCodec()
// 	cryptoAmino.RegisterAmino(cdc)
// 	cdc.RegisterInterface((*types.PrivValidator)(nil), nil)
// 	cdc.RegisterConcrete(&types.MockPV{}, "tendermint/MockPV", nil)

// 	var valList generator.ValList
// 	er := cdc.UnmarshalJSON(data, &valList)
// 	if er != nil {
// 		fmt.Printf("error: %v", er)
// 	}
// 	fmt.Printf("%+v", valList)
// }
