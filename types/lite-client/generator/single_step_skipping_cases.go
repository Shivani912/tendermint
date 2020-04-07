package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

// Validator Set

func caseSingleSkipOneBlock(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(
		valList.Validators[:3],
		valList.PrivVals[:3],
		1,
	)
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipFiveBlocks(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=7, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(
		valList.Validators[:3],
		valList.PrivVals[:3],
		5,
	)
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipValidatorSetChangesLessThanTrustLevel(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1 verifying signed header at height=7 while valset changes less than default trust level (1/3), should not expect error"

	copyValList := valList.Copy()
	initial, input, state, privVals := generateInitialAndInputSkipBlocks(
		copyValList.Validators[:4],
		copyValList.PrivVals[:4],
		3,
	)
	liteBlock, state, privVals := generateNextBlockWithNextValsUpdate(
		state,
		privVals,
		input[0].SignedHeader.Commit,
		copyValList.Validators[1:5],
		copyValList.PrivVals[1:5],
		thirdBlockTime.Add(30*time.Second),
	)
	liteBlock, state, _ = generateNextBlock(state, privVals, liteBlock.SignedHeader.Commit, thirdBlockTime.Add(35*time.Second))
	input[0] = liteBlock
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipValidatorSetChangesMoreThanTrustLevel(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=7 while valset changes more than default trust level (1/3), should expect error"

	copyValList := valList.Copy()
	initial, input, state, privVals := generateInitialAndInputSkipBlocks(
		copyValList.Validators[:4],
		copyValList.PrivVals[:4],
		3)
	liteBlock, state, privVals := generateNextBlockWithNextValsUpdate(
		state,
		privVals,
		input[0].SignedHeader.Commit,
		copyValList.Validators[3:7],
		copyValList.PrivVals[3:7],
		thirdBlockTime.Add(30*time.Second),
	)
	liteBlock, state, _ = generateNextBlock(state, privVals, liteBlock.SignedHeader.Commit, thirdBlockTime.Add(35*time.Second))
	input[0] = liteBlock
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

// Commit

func caseSingleSkipCommitOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3 where 1/3 vals dont sign, should expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(
		valList.Validators[:3],
		valList.PrivVals[:3],
		1,
	)
	input[0].SignedHeader.Commit.Signatures[0].BlockIDFlag = types.BlockIDFlagAbsent
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipCommitMoreThanTwoThirdsValsDidSign(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3 where more than two-thirds vals did sign, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(
		valList.Validators[:4],
		valList.PrivVals[:4],
		1,
	)
	input[0].SignedHeader.Commit.Signatures[0] = types.CommitSig{
		BlockIDFlag:      types.BlockIDFlagAbsent,
		ValidatorAddress: nil,
	}
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

// Header

func caseSingleSkipHeaderOutOfTrustingPeriod(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1 but is out of trusting period, verifying signed header at height=5, expects an error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(
		valList.Validators[:1],
		valList.PrivVals[:1],
		4,
	)
	initial.TrustingPeriod = 5 * time.Second

	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
