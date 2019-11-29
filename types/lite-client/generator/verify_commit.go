package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

func CaseVerifyCommitEmpty(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, empty commit, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit = nil
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongHeaderHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong header hash, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.Hash = []byte("wrong header hash!!")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongPartsHeaderCount(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Total, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Total += 5
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongPartsHeaderHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Hash, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Hash = []byte("wrong PartsHeader hash!!")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteType(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong vote type, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Type = types.SignedMsgType(types.PrevoteType)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteHeight(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong vote height, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Height -= 1
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteRound(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong vote round, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Round -= 1
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteBlockID(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong vote BlockID, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].BlockID.Hash = []byte("wrong hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteTimestamp(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong vote timestamp, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)

	wrongTimestamp, _ := time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")
	input[0].SignedHeader.Commit.Precommits[0].Timestamp = wrongTimestamp
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongVoteSignature(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong signature in vote, with error"
	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Signature = []byte("wrong address")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitOneThirdValsDontSign(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, one-third vals don't sign, expects error"
	initial, input, _, _ := GenerateGeneralCase(valList, 3)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitLessThanOneThirdValsDontSign(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, less than one-third vals don't sign, no error"
	initial, input, _, _ := GenerateGeneralCase(valList, 4)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)
	testCases.TC = append(testCases.TC, testCase)
}
