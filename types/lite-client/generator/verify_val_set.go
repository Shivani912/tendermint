package generator

import (
	"github.com/tendermint/tendermint/types"
)

func CaseVerifyValidatorSetOf1(testCases *TestCases, valList ValList) {

	// DONE: lets have a `testNameVerify` constant that we can define globally once and use for all these
	// instead of redefining this var each time

	description := "Case: one lite block to fetch, one validator in the set, expects no error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetOf8(testCases *TestCases, valList ValList) {

	description := "Case: one lite block to fetch, 8 validators in the set, expects no error"
	initial, input, _, _ := GenerateGeneralCase(valList, 8)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetOf128(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, 128 validators, no error"
	initial, input, _, _ := GenerateGeneralCase(valList, 128)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)
	testCases.TC = append(testCases.TC, testCase)

}

func CaseVerifyValidatorSetEmpty(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, empty validator set, expects error"
	initial, input, _, _ := GenerateGeneralCase(valList, 2)
	input[0].ValidatorSet = *types.NewValidatorSet(nil)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)

	testCases.TC = append(testCases.TC, testCase)

}

func CaseVerifyValidatorSetAddTwiceVals(testCases *TestCases, valList ValList) {

	description := "Case: two lite blocks, validator set reduces to half, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(valList, 2, 2, 0)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetRemoveHalfVals(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, validator set reduces to half, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 4, 0, 2)

	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)
	testCases.TC = append(testCases.TC, testCase)

}

func CaseVerifyValidatorSetChangesOneThird(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 1/3 validator set changes, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 3, 1, 1)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangesHalf(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 1/2 validator set changes, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 4, 2, 2)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangesTwoThirds(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 2/3 validator set changes, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 3, 2, 2)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangesFully(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, validator set changes completely, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 5, 5, 5)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangesLessThanOneThird(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, less than 1/3 validator set changes, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 4, 1, 1)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, more than 2/3 validator set changes, no error"

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, 4, 3, 3)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetWrongValidatorSet(testCases *TestCases, valList ValList) {

	var input []LiteBlock
	description := "Case: one lite block, wrong validator set, expects error"

	signedHeader, state, _ := GenerateFirstBlock(valList, 3, firstBlockTime)
	initial := GenerateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	wrongVals := valList.Validators[3:6]
	wrongPrivVals := valList.PrivVals[3:6]
	wrongValSet := types.NewValidatorSet(wrongVals)
	state.Validators = wrongValSet
	state.NextValidators = wrongValSet

	liteBlock, state := GenerateNextBlock(state, wrongPrivVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetReplaceValidator(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, replacing a validator in validator set, expects error"

	signedHeader, state, privVals := GenerateFirstBlock(copyValList, 3, firstBlockTime)
	initial := GenerateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	privVals[0] = copyValList.PrivVals[4]
	state.Validators.Validators[0] = copyValList.Validators[4]
	state.NextValidators = state.Validators

	liteBlock, state := GenerateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)

	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyValidatorSetChangeValidatorPower(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, changing a validator's power in validator set, expects error"

	signedHeader, state, privVals := GenerateFirstBlock(copyValList, 3, firstBlockTime)
	initial := GenerateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	// privVals[0] = copyValList.PrivVals[4]
	state.Validators.Validators[0].VotingPower += 1 //copyValList.Validators[4]
	state.NextValidators = state.Validators

	liteBlock, state := GenerateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)

	testCases.TC = append(testCases.TC, testCase)
}
