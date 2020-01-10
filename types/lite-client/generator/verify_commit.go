package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

var (
	str32byte = "----This is a 32-byte string----"
	str64byte = "----------This is a 64-byte long long long long string----------"
)

func caseVerifyCommitEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty commit, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit = nil
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongHeaderHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong header hash, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.Hash = []byte(str32byte)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongPartsHeaderCount(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Total, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Total += 5
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongPartsHeaderHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Hash, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Hash = []byte(str32byte)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteType(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote type, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Type = types.SignedMsgType(types.PrevoteType)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteHeight(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote height, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Height--
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteRound(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote round, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Round--
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteBlockID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote BlockID, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].BlockID.Hash = []byte(str32byte)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteTimestamp(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote timestamp, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)

	wrongTimestamp, _ := time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")
	input[0].SignedHeader.Commit.Precommits[0].Timestamp = wrongTimestamp
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitWrongVoteSignature(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong signature in vote, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Signature = []byte(str64byte)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, one-third vals don't sign, expects error"
	initial, input, _, _ := generateGeneralCase(valList, 3)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyCommitLessThanOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, less than one-third vals don't sign, no error"
	initial, input, _, _ := generateGeneralCase(valList, 4)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := generateTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
