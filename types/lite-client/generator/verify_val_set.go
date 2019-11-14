package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

func CaseVerifyValidatorSetOf1(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, one validator, no error")

	vals := valList.ValidatorSet.Validators[:1]
	privVal := valList.PrivVal[:1]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetOf8(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 8 validators, no error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetOf128(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, 128 validators, no error")

	vals := valList.ValidatorSet.Validators[:128]
	privVal := valList.PrivVal[:128]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func CaseVerifyValidatorSetEmpty(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, empty validator set, expects error")

	vals := valList.ValidatorSet.Validators[:8]
	privVal := valList.PrivVal[:8]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	testCase.Input[0].ValidatorSet = *types.NewValidatorSet(nil)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)

}

func CaseVerifyValidatorSetAddTwiceVals(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, validator set grows twice (2 + 2), no error")

	vals := valList.ValidatorSet.Validators[:2]
	privVal := valList.PrivVal[:2]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[2:4]
	newPrivVal := valList.PrivVal[2:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 0, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetRemoveHalfVals(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, validator set reduces to half, no error")

	vals := valList.ValidatorSet.Validators[:4]
	privVal := valList.PrivVal[:4]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	// newVals := valList.ValidatorSet.Validators[2:4]
	// newPrivVal := valList.PrivVal[2:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, nil, nil, 2, secondBlockTime)
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesOneThird(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 1/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[3:4]
	newPrivVal := valList.PrivVal[3:4]

	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 1, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesHalf(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 1/2 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:4]
	privVal := valList.PrivVal[:4]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[4:6]
	newPrivVal := valList.PrivVal[4:6]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 2, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesTwoThirds(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, 2/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[4:6]
	newPrivVal := valList.PrivVal[4:6]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 2, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesFully(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, validator set changes completely, no error")

	vals := valList.ValidatorSet.Validators[:5]
	privVal := valList.PrivVal[:5]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[5:10]
	newPrivVal := valList.PrivVal[5:10]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 5, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesLessThanOneThird(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, less than 1/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:4]
	privVal := valList.PrivVal[:4]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[4:5]
	newPrivVal := valList.PrivVal[4:5]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 1, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: two lite blocks, more than 2/3 validator set changes, no error")

	vals := valList.ValidatorSet.Validators[:4]
	privVal := valList.PrivVal[:4]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	newVals := valList.ValidatorSet.Validators[4:7]
	newPrivVal := valList.PrivVal[4:7]
	privVal = GenerateNextBlockWithNextValsUpdate(testCase, state, privVal, testCase.Initial.SignedHeader.Commit, newVals, newPrivVal, 3, secondBlockTime)

	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Input[0].SignedHeader.Commit, thirdBlockTime)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

// NOTE: commented on Lite Client PR already. This case should give error.
func CaseVerifyValidatorSetWrongProposer(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong proposer, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)

	vs := *state.Validators.Copy()
	state.Validators.IncrementProposerPriority(1)

	// fmt.Printf("\nvs: %+v, hash: %v \nvs: %+v, hash: %v", vs, vs.Hash(), state.Validators, state.Validators.Hash())

	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, vs, 3*time.Hour, now)

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyValidatorSetWrongValidatorSet(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong validator set, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
	vs := *state.Validators.Copy()

	wrongVals := valList.ValidatorSet.Validators[3:6]
	wrongPrivVal := valList.PrivVal[3:6]
	wrongValSet := types.NewValidatorSet(wrongVals)
	state.Validators = wrongValSet
	state.NextValidators = wrongValSet

	// fmt.Printf("\nvs: %+v, hash: %v \nvs: %+v, hash: %v", vs, vs.Hash(), state.Validators, state.Validators.Hash())

	GenerateNextBlock(state, testCase, wrongPrivVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)
	GenerateInitial(testCase, vs, 3*time.Hour, now)

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}
