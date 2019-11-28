package generator

func CaseVerifyCommitEmpty(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, empty commit, with error"

	testCase := GenerateGeneralTestCase(valList, 1, description, expectedOutputError)
	testCase.Input[0].SignedHeader.Commit = nil
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongHeaderHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong header hash, with error"
	testCase := GenerateGeneralTestCase(valList, 1, description, expectedOutputError)
	testCase.Input[0].SignedHeader.Commit.BlockID.Hash = []byte("wrong header hash!!")
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyCommitWrongPartsHeaderCount(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong PartsHeader.Total, with error"
	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Commit.BlockID.PartsHeader.Total += 5
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

// func CaseVerifyCommitWrongPartsHeaderHash(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong PartsHeader.Hash, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.BlockID.PartsHeader.Hash = []byte("wrong PartsHeader hash!!")
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteType(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong vote type, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].Type = types.SignedMsgType(types.PrevoteType)
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteHeight(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong vote height, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].Height -= 1
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteRound(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong vote round, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].Round -= 1
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteBlockID(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong vote BlockID, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].BlockID.Hash = []byte("wrong hash")
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteTimestamp(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong vote timestamp, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	wrongTimestamp, _ := time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")
// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].Timestamp = wrongTimestamp
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteSignature(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, wrong signature in vote, with error"
// 	GenerateGeneralTestCase(testCase, valList, 1, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0].Signature = []byte("wrong address")
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitWrongVoteInvalidSignature(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, invalid signature in vote, with error")

// 	vals := valList.ValidatorSet.Validators[:3]
// 	privVal := valList.PrivVal[:3]

// 	state := GenerateFirstBlock(testCase, vals, privVal, firstBlockTime)
// 	GenerateNextBlockWithDoubleSign(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, secondBlockTime)

// 	GenerateInitial(testCase, *state.Validators, 3*time.Hour, now)
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitOneThirdValsDontSign(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, one-third vals don't sign, expects error"
// 	GenerateGeneralTestCase(testCase, valList, 3, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0] = nil
// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }

// func CaseVerifyCommitLessThanOneThirdValsDontSign(testCases *TestCases, valList ValList) {
// 	var testCase *TestCase = &TestCase{}

// 	name := "verify"
// 	description := "Case: one lite block, less than one-third vals don't sign, no error"
// 	GenerateGeneralTestCase(testCase, valList, 4, name, description)

// 	testCase.Input[0].SignedHeader.Commit.Precommits[0] = nil

// 	GenerateExpectedOutput(testCase)
// 	testCases.TC = append(testCases.TC, *testCase)
// }
