package garbage

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

var now, _ = time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")

func GenerateTestCase() {

	var testCases *TestCases = &TestCases{}

	VerifyCase1(testCases)
	VerifyCase2(testCases)
	VerifyCase3(testCases)
	VerifyCase4(testCases)
	VerifyCase5(testCases)
	VerifyCase6(testCases)
	VerifyCase7(testCases)

	GenerateJSON(testCases)
}

func VerifyCase1(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 1: one lite block, one validator, no error")
	state, privVal := GenerateFirstBlock(testCase, 1, 10)

	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func VerifyCase2(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 2: two lite blocks, validator set grows 2x (1 + 1), no error")
	state, privVal := GenerateFirstBlock(testCase, 1, 10)
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, 1, 10)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func VerifyCase3(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 3: two lite blocks, 1/3 validator set changes, no error")
	state, privVal := GenerateFirstBlock(testCase, 3, 10)
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, 1, 10)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func VerifyCase4(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 4: one lite block, 8 validators, no error")
	state, privVal := GenerateFirstBlock(testCase, 8, 90)
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func VerifyCase5(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 5: one lite block, 16 validators, no error")
	state, privVal := GenerateFirstBlock(testCase, 16, 900)
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func VerifyCase6(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 6: one lite block, 32 validators, no error")
	state, privVal := GenerateFirstBlock(testCase, 32, 900)
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func VerifyCase7(testCases *TestCases) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case 7: empty validator set, with error")
	state, privVal := GenerateFirstBlock(testCase, 3, 100)
	privVal = GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	testCase.Input[0].ValidatorSet = *types.NewValidatorSet(nil)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}
