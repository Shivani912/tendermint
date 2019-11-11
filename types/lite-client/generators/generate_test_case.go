package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

var now, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")

func GenerateTestCase(jsonValList string) {

	var testCases *TestCases = &TestCases{}
	valList := GetValList(jsonValList)

	CaseVerifyValidatorSetOf1(testCases, valList)
	CaseVerifyValidatorSetOf8(testCases, valList)
	CaseVerifyValidatorSetOf128(testCases, valList)
	CaseVerifyValidatorSetEmpty(testCases, valList)

	CaseVerifyValidatorSetAddTwiceVals(testCases, valList)

	CaseVerifyValidatorSetChangesOneThird(testCases, valList)
	valList = GetValList(jsonValList)
	CaseVerifyValidatorSetChangesHalf(testCases, valList)

	GenerateJSON(testCases)
}

func CaseVerifyValidatorSetOf1(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, one validator, no error")

	vals := valList.ValidatorSet.Validators[:1]
	privVal := valList.PrivVal[:1]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetAddTwiceVals(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, validator set grows twice (2 + 2), no error")

	vals := valList.ValidatorSet.Validators[:2]
	privVal := valList.PrivVal[:2]

	state := GenerateFirstBlock(testCase, vals, privVal, now)

	newVals := valList.ValidatorSet.Validators[2:4]
	newPrivVal := valList.PrivVal[2:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 0)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, now.Add(time.Second*5))
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesOneThird(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 1/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)

	newVals := valList.ValidatorSet.Validators[3:4]
	newPrivVal := valList.PrivVal[3:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 1)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, now.Add(time.Second*5))
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesHalf(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 1/2 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:4]
	privVal := valList.PrivVal[:4]

	state := GenerateFirstBlock(testCase, vals, privVal, now)

	newVals := valList.ValidatorSet.Validators[4:6]
	newPrivVal := valList.PrivVal[4:6]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 2)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	// fmt.Printf("vals: \n %v \n priv vals: \n %v", state.Validators, privVal)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, now.Add(time.Second*5))
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetOf8(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 8 validators, no error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetEmpty(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, empty validator set, expects error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].ValidatorSet = *types.NewValidatorSet(nil)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func CaseVerifyValidatorSetOf128(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 128 validators, no error")

	vals := valList.ValidatorSet.Validators[:128]
	privVal := valList.PrivVal[:128]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}
