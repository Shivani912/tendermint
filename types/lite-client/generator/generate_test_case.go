package generator

// GenerateVerifyTestCases is a wrapper function around all the test case specific functions.
// It calls these functions that essentially make the test cases and store them under TestCases data structure
// These cases are categorized according to the data structure it is trying to test (e.g. Validator set, Commit, etc...)
func GenerateVerifyTestCases(jsonValList string) {

	valList := GetValList(jsonValList)

	// Verify - ValidatorSet
	testBatch := newBatch("verify-validator set")
	caseVerifyValidatorSetOf1(testBatch, valList)
	caseVerifyValidatorSetOf8(testBatch, valList)
	caseVerifyValidatorSetOf128(testBatch, valList)
	caseVerifyValidatorSetEmpty(testBatch, valList)

	caseVerifyValidatorSetAddTwiceVals(testBatch, valList)
	caseVerifyValidatorSetRemoveHalfVals(testBatch, valList)

	caseVerifyValidatorSetChangesOneThird(testBatch, valList)
	caseVerifyValidatorSetChangesHalf(testBatch, valList)
	caseVerifyValidatorSetChangesTwoThirds(testBatch, valList)
	caseVerifyValidatorSetChangesFully(testBatch, valList)
	caseVerifyValidatorSetChangesLessThanOneThird(testBatch, valList)
	caseVerifyValidatorSetChangesMoreThanTwoThirds(testBatch, valList)
	caseVerifyValidatorSetWrongValidatorSet(testBatch, valList)
	caseVerifyValidatorSetReplaceValidator(testBatch, valList)
	caseVerifyValidatorSetChangeValidatorPower(testBatch, valList)

	generateJSON(testBatch, "./tests/json/val_set_tests.json")

	// Verify - Commit
	testBatch = newBatch("verify-commit")
	caseVerifyCommitEmpty(testBatch, valList)
	caseVerifyCommitWrongHeaderHash(testBatch, valList)
	caseVerifyCommitWrongPartsHeaderCount(testBatch, valList)
	caseVerifyCommitWrongPartsHeaderHash(testBatch, valList)
	caseVerifyCommitWrongVoteType(testBatch, valList)
	caseVerifyCommitWrongVoteHeight(testBatch, valList)
	caseVerifyCommitWrongVoteRound(testBatch, valList)
	caseVerifyCommitWrongVoteBlockID(testBatch, valList)
	caseVerifyCommitWrongVoteTimestamp(testBatch, valList)
	caseVerifyCommitWrongVoteSignature(testBatch, valList)

	// TODO: more cases
	// - commits from wrong validators
	// We need to come back to this after the commit structure changes
	caseVerifyCommitOneThirdValsDontSign(testBatch, valList)         // error
	caseVerifyCommitLessThanOneThirdValsDontSign(testBatch, valList) // not an error

	generateJSON(testBatch, "./tests/json/commit_tests.json")

	// Verify - Header
	testBatch = newBatch("verify-header")
	caseVerifyHeaderEmpty(testBatch, valList)
	caseVerifyHeaderWrongLastCommitHash(testBatch, valList)
	caseVerifyHeaderWrongLastResultsHash(testBatch, valList)
	caseVerifyHeaderWrongLastBlockID(testBatch, valList)
	caseVerifyHeaderWrongChainID(testBatch, valList)
	caseVerifyHeaderWrongHeight(testBatch, valList)
	caseVerifyHeaderWrongTimestamp(testBatch, valList)
	caseVerifyHeaderWrongValSetHash(testBatch, valList)
	caseVerifyHeaderWrongNextValSetHash(testBatch, valList)

	generateJSON(testBatch, "./tests/json/header_tests.json")
}

func newBatch(name string) *TestBatch {
	return &TestBatch{
		BatchName: name,
	}
}
