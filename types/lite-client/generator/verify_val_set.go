package generator

import (
	"github.com/tendermint/tendermint/types"
)

func caseVerifyValidatorSetOf1(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block to fetch, one validator in the set, expects no error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetOf8(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block to fetch, 8 validators in the set, expects no error"
	initial, input, _, _ := generateGeneralCase(valList, 8)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetOf128(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, 128 validators, no error"
	initial, input, _, _ := generateGeneralCase(valList, 128)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseVerifyValidatorSetEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty validator set, expects error"
	initial, input, _, _ := generateGeneralCase(valList, 2)
	input[0].ValidatorSet = *types.NewValidatorSet(nil)
	testCase := generateTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseVerifyValidatorSetAddTwiceVals(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, validator set reduces to half, no error"

	initial, input, _, _ := generateNextValsUpdateCase(valList, 2, 2, 0)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetRemoveHalfVals(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, validator set reduces to half, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 4, 0, 2)

	testCase := generateTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseVerifyValidatorSetChangesOneThird(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 1/3 validator set changes, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 3, 1, 1)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangesHalf(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 1/2 validator set changes, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 4, 2, 2)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangesTwoThirds(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, 2/3 validator set changes, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 3, 2, 2)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangesFully(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, validator set changes completely, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 5, 5, 5)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangesLessThanOneThird(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, less than 1/3 validator set changes, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 4, 1, 1)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangesMoreThanTwoThirds(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	description := "Case: two lite blocks, more than 2/3 validator set changes, no error"

	initial, input, _, _ := generateNextValsUpdateCase(copyValList, 4, 3, 3)
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetWrongValidatorSet(testBatch *TestBatch, valList ValList) {

	var input []LiteBlock
	description := "Case: one lite block, wrong validator set, expects error"

	signedHeader, state, _ := generateFirstBlock(valList, 3, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	wrongVals := valList.Validators[3:6]
	wrongPrivVals := valList.PrivVals[3:6]
	wrongValSet := types.NewValidatorSet(wrongVals)
	state.Validators = wrongValSet
	state.NextValidators = wrongValSet

	liteBlock, state := generateNextBlock(state, wrongPrivVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := generateTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetReplaceValidator(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, replacing a validator in validator set, expects error"

	signedHeader, state, privVals := generateFirstBlock(copyValList, 3, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	privVals[0] = copyValList.PrivVals[4]
	state.Validators.Validators[0] = copyValList.Validators[4]
	state.NextValidators = state.Validators

	liteBlock, state := generateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := generateTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyValidatorSetChangeValidatorPower(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, changing a validator's power in validator set, expects error"

	signedHeader, state, privVals := generateFirstBlock(copyValList, 3, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	state.Validators.Validators[0].VotingPower++
	state.NextValidators = state.Validators

	liteBlock, state := generateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := generateTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
