package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

var (
	str32byte = "----This is a 32-byte string----"
	str64byte = []byte{206, 129, 9, 176, 142, 141, 188, 30, 197, 158, 80, 135, 172, 5, 239, 44, 219, 46, 60, 239, 9, 65, 151, 236, 221, 44, 72, 253, 191, 95, 20, 67, 175, 2, 133, 74, 3, 84, 20, 60, 142, 1, 0, 75, 129, 148, 2, 206, 180, 49, 223, 47, 41, 189, 149, 230, 247, 16, 48, 228, 39, 91, 154, 6}

	//"----------This is a 64-byte long long long long string----------"
)

// HEADER - BEGIN

func caseSingleSeqHeaderEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header = nil
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongLastCommitHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last commit hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastCommitHash = []byte("wrong hash")
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongLastResultsHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last results hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastResultsHash = []byte("wrong hash")
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
func caseSingleSeqHeaderWrongLastBlockID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last block ID in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastBlockID.Hash = []byte("wrong hash")
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongChainID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong chain ID in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ChainID = "wrong-chain-id"
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongHeight(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong height in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Height++
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongTimestamp(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong timestamp in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Time = secondBlockTime.Add(1 * time.Minute)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongValSetHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong val set hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ValidatorsHash = []byte("wrong validator set hash")
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqHeaderWrongNextValSetHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong next val set hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.NextValidatorsHash = []byte("wrong next validator set hash")
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

// COMMIT - BEGIN
func caseSingleSeqCommitEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty commit, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit = nil
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongHeaderHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong header hash, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.Hash = []byte(str32byte)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongPartsHeaderCount(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Total, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Total += 5
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongPartsHeaderHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong PartsHeader.Hash, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.BlockID.PartsHeader.Hash = []byte(str32byte)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteType(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote type, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Type = types.SignedMsgType(types.PrevoteType)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteHeight(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote height, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Height--
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteRound(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote round, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].Round--
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteBlockID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote BlockID, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Commit.Precommits[0].BlockID.Hash = []byte(str32byte)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteTimestamp(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong vote timestamp, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)

	wrongTimestamp, _ := time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")
	input[0].SignedHeader.Commit.Precommits[0].Timestamp = wrongTimestamp
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitWrongVoteSignature(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong signature in vote, with error"
	initial, input, _, _ := generateGeneralCase(valList, 1)
	// fmt.Println("", input[0].SignedHeader.Commit.Precommits[0].Signature)
	input[0].SignedHeader.Commit.Precommits[0].Signature = []byte(str64byte)
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitOneThirdValsDontSign(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, one-third vals don't sign, expects error"
	initial, input, _, _ := generateGeneralCase(valList, 3)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := makeTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqCommitMoreThanTwoThirdsValsDidSign(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, less than one-third vals don't sign, no error"
	initial, input, _, _ := generateGeneralCase(valList, 4)
	input[0].SignedHeader.Commit.Precommits[0] = nil
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

// VALIDATOR SET - BEGIN

func caseSingleSeqValidatorSetOf1(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block to fetch, one validator in the set, expects no error"
	testCase := generateAndMakeGeneralTestCase(description, valList, 1, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetOf8(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block to fetch, 8 validators in the set, expects no error"
	testCase := generateAndMakeGeneralTestCase(description, valList, 8, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetOf128(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, 128 validators, no error"
	testCase := generateAndMakeGeneralTestCase(description, valList, 128, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseSingleSeqValidatorSetEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty validator set, expects error"
	initial, input, _, _ := generateGeneralCase(valList, 2)
	input[0].ValidatorSet = *types.NewValidatorSet(nil)
	testCase := makeTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseSingleSeqValidatorSetAddTwiceVals(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, validator set reduces to half, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 2, 2, 0, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetRemoveHalfVals(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, validator set reduces to half, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 4, 0, 2, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)

}

func caseSingleSeqValidatorSetChangesOneThird(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, 1/3 validator set changes, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 3, 1, 1, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangesHalf(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, 1/2 validator set changes, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 4, 2, 2, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangesTwoThirds(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, 2/3 validator set changes, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 3, 2, 2, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangesFully(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, validator set changes completely, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 5, 5, 5, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangesLessThanOneThird(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, less than 1/3 validator set changes, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 4, 1, 1, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangesMoreThanTwoThirds(testBatch *TestBatch, valList ValList) {

	description := "Case: two lite blocks, more than 2/3 validator set changes, no error"
	testCase := generateAndMakeNextValsUpdateTestCase(description, valList, 4, 3, 3, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetWrongValidatorSet(testBatch *TestBatch, valList ValList) {

	var input []LiteBlock
	description := "Case: one lite block, wrong validator set, expects error"

	signedHeader, state, _ := generateFirstBlock(valList, 3, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	wrongVals := valList.Validators[3:6]
	wrongPrivVals := valList.PrivVals[3:6]
	wrongValSet := types.NewValidatorSet(wrongVals)
	state.Validators = wrongValSet
	state.NextValidators = wrongValSet

	liteBlock, state, _ := generateNextBlock(state, wrongPrivVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := makeTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetFaultySigner(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, faulty signer (not present in validator set), expects error"

	signedHeader, state, privVals := generateFirstBlock(copyValList, 4, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	liteBlock, state, _ := generateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)

	liteBlock.ValidatorSet = *types.NewValidatorSet(copyValList.Validators[:3])
	liteBlock.SignedHeader.Header.ValidatorsHash = liteBlock.ValidatorSet.Hash()
	liteBlock.SignedHeader.Commit.BlockID.Hash = liteBlock.SignedHeader.Header.Hash()
	liteBlock.SignedHeader.Commit.Signatures = liteBlock.SignedHeader.Commit.Signatures[1:4]

	initial.NextValidatorSet = liteBlock.ValidatorSet

	input = append(input, liteBlock)
	testCase := makeTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseSingleSeqValidatorSetChangeValidatorPower(testBatch *TestBatch, valList ValList) {

	copyValList := valList.Copy()
	var input []LiteBlock
	description := "Case: one lite block, changing a validator's power in validator set, expects error"

	signedHeader, state, privVals := generateFirstBlock(copyValList, 3, firstBlockTime)
	initial := generateInitial(signedHeader, *state.NextValidators, trustingPeriod, now)

	state.Validators.Validators[0].VotingPower++
	state.NextValidators = state.Validators

	liteBlock, state, _ := generateNextBlock(state, privVals, initial.SignedHeader.Commit, secondBlockTime)
	input = append(input, liteBlock)
	testCase := makeTestCase(description, initial, input, expectedOutputError)

	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
