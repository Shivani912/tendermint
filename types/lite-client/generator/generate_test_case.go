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

	// DONE: let's avoid this reloading by copying in the CaseVerifyXXX where necessary
	CaseVerifyValidatorSetChangesOneThird(testCases, valList)
	CaseVerifyValidatorSetChangesHalf(testCases, valList)
	CaseVerifyValidatorSetChangesTwoThirds(testCases, valList)
	CaseVerifyValidatorSetChangesFully(testCases, valList)
	CaseVerifyValidatorSetChangesLessThanOneThird(testCases, valList)
	CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases, valList)

	// TODO: how about some variations on the wrong validator set:
	// - replace a validator
	// - change a validators power
	// - different validator set
	CaseVerifyValidatorSetWrongValidatorSet(testCases, valList)

	// Verify - Commit
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

	// Verify - Header
	CaseVerifyHeaderEmpty(testCases, valList)
	CaseVerifyHeaderWrongLastCommitHash(testCases, valList)
	CaseVerifyHeaderWrongLastResultsHash(testCases, valList)
	CaseVerifyHeaderWrongLastBlockID(testCases, valList)
	CaseVerifyHeaderWrongChainID(testCases, valList)
	CaseVerifyHeaderWrongHeight(testCases, valList)
	CaseVerifyHeaderWrongTimestamp(testCases, valList)
	CaseVerifyHeaderWrongValSetHash(testCases, valList)
	CaseVerifyHeaderWrongNextValSetHash(testCases, valList)

	file := "./tests/json/test_lite_client.json"
	GenerateJSON(testCases, file)
}
