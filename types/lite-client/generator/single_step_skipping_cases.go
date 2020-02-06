package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

func caseSingleSkipOneBlock(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(valList, 3, 1)
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipFiveBlocks(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=7, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(valList, 3, 5)
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipValidatorSetChangesLessThanTrustLevel(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1 verifying signed header at height=7 while valset changes less than default trust level (1/3), should not expect error"

	copyValList := valList.Copy()
	initial, input, state, privVals := generateInitialAndInputSkipBlocks(copyValList, 4, 3)
	liteBlock, state, privVals := generateNextBlockWithNextValsUpdate(state, privVals, input[0].SignedHeader.Commit, copyValList, 10, 11, 1, thirdBlockTime.Add(30*time.Second))
	liteBlock, state, _ = generateNextBlock(state, privVals, liteBlock.SignedHeader.Commit, thirdBlockTime.Add(35*time.Second))
	input[0] = liteBlock
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipValidatorSetChangesMoreThanTrustLevel(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=7 while valset changes more than default trust level (1/3), should expect error"

	copyValList := valList.Copy()
	initial, input, state, privVals := generateInitialAndInputSkipBlocks(copyValList, 4, 3)
	liteBlock, state, privVals := generateNextBlockWithNextValsUpdate(state, privVals, input[0].SignedHeader.Commit, copyValList, 10, 13, 3, thirdBlockTime.Add(30*time.Second))
	liteBlock, state, _ = generateNextBlock(state, privVals, liteBlock.SignedHeader.Commit, thirdBlockTime.Add(35*time.Second))
	input[0] = liteBlock
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipCommitOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3 where 1/3 vals dont sign, should expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(valList, 3, 1)
	input[0].SignedHeader.Commit.Signatures[0].BlockIDFlag = types.BlockIDFlagAbsent
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSkipCommitLessThanOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, verifying signed header at height=3 where less than 1/3 vals dont sign, should not expect error"

	initial, input, _, _ := generateInitialAndInputSkipBlocks(valList, 4, 1)
	input[0].SignedHeader.Commit.Signatures[0].BlockIDFlag = types.BlockIDFlagAbsent
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
