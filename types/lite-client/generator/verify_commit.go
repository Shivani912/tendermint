package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

func CaseVerifyCommitEmpty(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, empty commit, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit = nil
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongHeaderHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong header hash, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.BlockID.Hash = []byte("wrong header hash!!")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongPartsHeaderCount(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong PartsHeader.Total, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.BlockID.PartsHeader.Total += 5
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongPartsHeaderHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong PartsHeader.Hash, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.BlockID.PartsHeader.Hash = []byte("wrong PartsHeader hash!!")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongVoteType(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong vote type, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.Precommits[0].Type = types.SignedMsgType(types.PrevoteType)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongVoteHeight(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong vote height, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.Precommits[0].Height -= 1
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyCommitWrongVoteRound(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, wrong vote round, with error")

	vals := valList.ValidatorSet.Validators[:3]
	privVal := valList.PrivVal[:3]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now.Add(time.Second*10))
	testCase.Input[0].SignedHeader.Commit.Precommits[0].Round -= 1
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}
