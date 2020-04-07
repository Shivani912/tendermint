package generator

// GenerateSingleStepSequentialCases creates three json files each for validator set, commit and header cases
// These cases are categorized according to the data structure it is trying to test (e.g. Validator set, Commit, etc...)
// It produces cases that test the single step sequential verification
// what this means is, given a trusted state and height can the lite node verify the next block height?
func GenerateSingleStepSequentialCases(jsonValList string) {

	valList := GetValList(jsonValList)

	// ValidatorSet
	testBatch := newBatch("Single Step Sequential-validator set")
	caseSingleSeqValidatorSetOf1(testBatch, valList)
	caseSingleSeqValidatorSetOf8(testBatch, valList)
	caseSingleSeqValidatorSetOf128(testBatch, valList)

	caseSingleSeqValidatorSetAddTwiceVals(testBatch, valList)
	caseSingleSeqValidatorSetRemoveHalfVals(testBatch, valList)

	caseSingleSeqValidatorSetChangesOneThird(testBatch, valList)
	caseSingleSeqValidatorSetChangesHalf(testBatch, valList)
	caseSingleSeqValidatorSetChangesTwoThirds(testBatch, valList)
	caseSingleSeqValidatorSetChangesFully(testBatch, valList)
	caseSingleSeqValidatorSetChangesLessThanOneThird(testBatch, valList)
	caseSingleSeqValidatorSetChangesMoreThanTwoThirds(testBatch, valList)
	caseSingleSeqValidatorSetWrongValidatorSet(testBatch, valList)
	caseSingleSeqValidatorSetFaultySigner(testBatch, valList)
	caseSingleSeqValidatorSetChangeValidatorPower(testBatch, valList)

	generateJSON(testBatch, "./tests/json/single_step_sequential/val_set_tests.json")

	// Commit
	testBatch = newBatch("Single Step Sequential-commit")
	caseSingleSeqCommitWrongHeaderHash(testBatch, valList)
	caseSingleSeqCommitWrongVoteHeight(testBatch, valList)
	caseSingleSeqCommitWrongVoteSignature(testBatch, valList)

	caseSingleSeqCommitOneThirdValsDontSign(testBatch, valList)         // error
	caseSingleSeqCommitMoreThanTwoThirdsValsDidSign(testBatch, valList) // not an error

	generateJSON(testBatch, "./tests/json/single_step_sequential/commit_tests.json")

	// Header
	testBatch = newBatch("Single Step Sequential-header")
	caseSingleSeqHeaderWrongHeaderSignature(testBatch, valList)
	caseSingleSeqHeaderWrongLastBlockID(testBatch, valList)
	caseSingleSeqHeaderWrongChainID(testBatch, valList)
	caseSingleSeqHeaderWrongHeight(testBatch, valList)
	caseSingleSeqHeaderWrongTimestamp(testBatch, valList)
	caseSingleSeqHeaderWrongValSetHash(testBatch, valList)
	caseSingleSeqHeaderWrongNextValSetHash(testBatch, valList)

	generateJSON(testBatch, "./tests/json/single_step_sequential/header_tests.json")
}

func newBatch(name string) *TestBatch {
	return &TestBatch{
		BatchName: name,
	}
}

// GenerateSingleStepSkippingCases creates three json files each for validator set, commit and header cases
// These cases test the single step skipping verification
// which means, given a trusted height and state can the lite node jump to a certain block height?
func GenerateSingleStepSkippingCases(jsonValList string) {

	valList := GetValList(jsonValList)

	// ValidatorSet
	testBatch := newBatch("Single Step Skipping-validator set")
	caseSingleSkipOneBlock(testBatch, valList)
	caseSingleSkipFiveBlocks(testBatch, valList)
	caseSingleSkipValidatorSetChangesLessThanTrustLevel(testBatch, valList)
	caseSingleSkipValidatorSetChangesMoreThanTrustLevel(testBatch, valList)

	generateJSON(testBatch, "./tests/json/single_step_skipping/val_set_tests.json")

	// Commit
	testBatch = newBatch("Single Step Skipping-commit")

	caseSingleSkipCommitOneThirdValsDontSign(testBatch, valList)         // error
	caseSingleSkipCommitMoreThanTwoThirdsValsDidSign(testBatch, valList) // not an error

	generateJSON(testBatch, "./tests/json/single_step_skipping/commit_tests.json")

	// Header
	testBatch = newBatch("Single Step Skipping-header")

	caseSingleSkipHeaderOutOfTrustingPeriod(testBatch, valList)

	generateJSON(testBatch, "./tests/json/single_step_skipping/header_tests.json")
}

func GenerateManyHeaderBisectionCases(jsonValList string) {

	valList := GetValList(jsonValList)

	// Single-peer
	caseBisectionHappyPath(valList)
	caseBisectionWorstCase(valList)
	caseBisectionInvalidValidatorSet(valList)
	caseBisectionNotEnoughCommits(valList)
	caseBisectionHeaderOutOfTrustingPeriod(valList)

	// Multi-peer
	caseBisectionConflictingValidCommitsFromTheOnlyWitness(valList)
	caseBisectionConflictingValidCommitsFromOneOfTheWitnesses(valList)
	caseBisectionConflictingHeaders(valList)
}
