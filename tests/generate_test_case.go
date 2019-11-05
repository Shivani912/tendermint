package tests

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

var now, _ = time.Parse(time.RFC3339, "2019-01-02T15:04:05Z")

func GenerateTestCase() {

	var testCases *TestCases = &TestCases{}
	valList := GetValList()

	CaseVerifyValidatorSetOf1(testCases, valList)
	CaseVerifyValidatorSetAddTwiceVals(testCases, valList)
	CaseVerifyValidatorSetOf8(testCases, valList)
	CaseVerifyValidatorSetOf128(testCases, valList)
	CaseVerifyValidatorSetEmpty(testCases, valList)
	CaseVerifyValidatorSetChangesOneThird(testCases, valList)

	GenerateJSON(testCases)
}

func CaseVerifyValidatorSetOf1(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, one validator, no error")

	vals := valList.ValidatorSet.Validators[:1]
	privVal := valList.PrivVal[:1]

	state := GenerateFirstBlock(testCase, vals, privVal)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetAddTwiceVals(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, validator set grows twice (2 + 2), no error")

	vals := valList.ValidatorSet.Validators[:2]
	privVal := valList.PrivVal[:2]

	state := GenerateFirstBlock(testCase, vals, privVal)

	newVals := valList.ValidatorSet.Validators[2:4]
	newPrivVal := valList.PrivVal[2:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesOneThird(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 1/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal)

	// newVals := valList.ValidatorSet.Validators[3:4]
	// newPrivVal := valList.PrivVal[3:4]

	state.Validators.Validators[2] = valList.ValidatorSet.Validators[3]
	privVal[2] = valList.PrivVal[3]

	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)

	// privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, 1, 10)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetOf8(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 8 validators, no error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetEmpty(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, empty validator set, expects error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	testCase.Input[0].ValidatorSet = *types.NewValidatorSet(nil)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func CaseVerifyValidatorSetOf128(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 128 validators, no error")

	vals := valList.ValidatorSet.Validators[:128]
	privVal := valList.PrivVal[:128]

	state := GenerateFirstBlock(testCase, vals, privVal)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, time.Now())
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}
