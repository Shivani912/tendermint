package generator

func GenerateTestCases(jsonValList string) {

	var testCases *TestCases = &TestCases{}
	valList := GetValList(jsonValList)

	// Verify - ValidatorSet
	CaseVerifyValidatorSetOf1(testCases, valList)
	CaseVerifyValidatorSetOf8(testCases, valList)
	// CaseVerifyValidatorSetOf128(testCases, valList)
	// CaseVerifyValidatorSetEmpty(testCases, valList)

	// CaseVerifyValidatorSetAddTwiceVals(testCases, valList)
	// CaseVerifyValidatorSetRemoveHalfVals(testCases, valList)

	// // TODO: let's avoid this reloading by copying in the CaseVerifyXXX where necessary
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesOneThird(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesHalf(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesTwoThirds(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesFully(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesLessThanOneThird(testCases, valList)
	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetChangesMoreThanTwoThirds(testCases, valList)

	// valList = GetValList(jsonValList)
	// CaseVerifyValidatorSetWrongProposer(testCases, valList)

	// // TODO: how about some variations on the wrong validator set:
	// // - replace a validator
	// // - change a validators power
	// // - different validator set
	// CaseVerifyValidatorSetWrongValidatorSet(testCases, valList)

	// // Verify - Commit
	// CaseVerifyCommitEmpty(testCases, valList)
	// CaseVerifyCommitWrongHeaderHash(testCases, valList)
	// CaseVerifyCommitWrongPartsHeaderCount(testCases, valList)
	// CaseVerifyCommitWrongPartsHeaderHash(testCases, valList)
	// CaseVerifyCommitWrongVoteType(testCases, valList)
	// CaseVerifyCommitWrongVoteHeight(testCases, valList)
	// CaseVerifyCommitWrongVoteRound(testCases, valList)
	// CaseVerifyCommitWrongVoteBlockID(testCases, valList)
	// CaseVerifyCommitWrongVoteTimestamp(testCases, valList)
	// CaseVerifyCommitWrongVoteSignature(testCases, valList)
	// CaseVerifyCommitWrongVoteInvalidSignature(testCases, valList)
	// valList = GetValList(jsonValList)

	// // TODO: more cases
	// // - commits from wrong validators
	// // We need to come back to this after the commit structure changes
	// CaseVerifyCommitOneThirdValsDontSign(testCases, valList)         // error
	// CaseVerifyCommitLessThanOneThirdValsDontSign(testCases, valList) // not an error

	// // Verify - Header
	// CaseVerifyHeaderEmpty(testCases, valList)
	// CaseVerifyHeaderWrongLastCommitHash(testCases, valList)
	// CaseVerifyHeaderWrongLastResultsHash(testCases, valList)
	// CaseVerifyHeaderWrongLastBlockID(testCases, valList)
	// CaseVerifyHeaderWrongDataHash(testCases, valList)
	// CaseVerifyHeaderWrongChainID(testCases, valList)
	// CaseVerifyHeaderWrongVersion(testCases, valList)
	// CaseVerifyHeaderWrongHeight(testCases, valList)
	// CaseVerifyHeaderWrongTimestamp(testCases, valList)
	// CaseVerifyHeaderWrongNumTxs(testCases, valList)
	// CaseVerifyHeaderWrongTotalTxs(testCases, valList)
	// CaseVerifyHeaderWrongValSetHash(testCases, valList)
	// CaseVerifyHeaderWrongNextValSetHash(testCases, valList)
	// CaseVerifyHeaderWrongAppHash(testCases, valList)
	// CaseVerifyHeaderWrongEvidenceHash(testCases, valList)
	// CaseVerifyHeaderWrongProposerAddress(testCases, valList)

	file := "./tests/json/test_lite_client.json"
	GenerateJSON(testCases, file)
}
