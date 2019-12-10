package generator

func GenerateTestCases(jsonValList string) {

	var testCases *TestCases = &TestCases{}
	valList := GetValList(jsonValList)

	// Verify - ValidatorSet
	CaseVerifyValidatorSetOf1(testCases, valList)
	CaseVerifyValidatorSetOf8(testCases, valList)
	CaseVerifyValidatorSetOf128(testCases, valList)
	CaseVerifyValidatorSetEmpty(testCases, valList)

	CaseVerifyValidatorSetAddTwiceVals(testCases, valList)
	CaseVerifyValidatorSetRemoveHalfVals(testCases, valList)

	CaseVerifyValidatorSetChangesOneThird(testCases, valList)
	CaseVerifyValidatorSetChangesHalf(testCases, valList)
	CaseVerifyValidatorSetChangesTwoThirds(testCases, valList)
	CaseVerifyValidatorSetChangesFully(testCases, valList)
	CaseVerifyValidatorSetChangesLessThanOneThird(testCases, valList)
	CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases, valList)
	CaseVerifyValidatorSetWrongValidatorSet(testCases, valList)
	CaseVerifyValidatorSetReplaceValidator(testCases, valList)
	CaseVerifyValidatorSetChangeValidatorPower(testCases, valList)

	GenerateJSON(testCases, "./tests/json/val_set_tests.json")

	// Verify - Commit
	testCases = &TestCases{}
	CaseVerifyCommitEmpty(testCases, valList)
	CaseVerifyCommitWrongHeaderHash(testCases, valList)
	CaseVerifyCommitWrongPartsHeaderCount(testCases, valList)
	CaseVerifyCommitWrongPartsHeaderHash(testCases, valList)
	CaseVerifyCommitWrongVoteType(testCases, valList)
	CaseVerifyCommitWrongVoteHeight(testCases, valList)
	CaseVerifyCommitWrongVoteRound(testCases, valList)
	CaseVerifyCommitWrongVoteBlockID(testCases, valList)
	CaseVerifyCommitWrongVoteTimestamp(testCases, valList)
	CaseVerifyCommitWrongVoteSignature(testCases, valList)

	// TODO: more cases
	// - commits from wrong validators
	// We need to come back to this after the commit structure changes
	CaseVerifyCommitOneThirdValsDontSign(testCases, valList)         // error
	CaseVerifyCommitLessThanOneThirdValsDontSign(testCases, valList) // not an error

	GenerateJSON(testCases, "./tests/json/commit_tests.json")

	// Verify - Header
	testCases = &TestCases{}
	CaseVerifyHeaderEmpty(testCases, valList)
	CaseVerifyHeaderWrongLastCommitHash(testCases, valList)
	CaseVerifyHeaderWrongLastResultsHash(testCases, valList)
	CaseVerifyHeaderWrongLastBlockID(testCases, valList)
	CaseVerifyHeaderWrongChainID(testCases, valList)
	CaseVerifyHeaderWrongHeight(testCases, valList)
	CaseVerifyHeaderWrongTimestamp(testCases, valList)
	CaseVerifyHeaderWrongValSetHash(testCases, valList)
	CaseVerifyHeaderWrongNextValSetHash(testCases, valList)

	GenerateJSON(testCases, "./tests/json/header_tests.json")
}
