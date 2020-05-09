package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

const MULTI_PEER_BISECTION_PATH = "./tests/json/bisection/multi_peer/"

func caseBisectionConflictingValidCommitsFromTheOnlyWitness(valList ValList) {
	description := "Case: Trusted height=1, found conflicting valid commit at height=11 from the only witness available, should expect error"
	primaryValSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	alternativeValSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	last := len(alternativeValSetChanges) - 1
	alternativeValSetChanges[last].Validators = valList.Validators[6:8]
	alternativeValSetChanges[last].PrivVals = valList.PrivVals[6:8]
	testBisection, _, _, _, _ := generateMultiPeerBisectionCase(
		description,
		primaryValSetChanges,
		alternativeValSetChanges,
		2,
		expectedOutputError,
	)

	file := MULTI_PEER_BISECTION_PATH + "conflicting_valid_commits_from_the_only_witness.json"
	testBisection.genJSON(file)
}

func caseBisectionConflictingValidCommitsFromOneOfTheWitnesses(valList ValList) {
	description := "Case: Trusted height=1, found conflicting valid commit at height=11 from only one of the two witnesses, should not expect error"
	primaryValSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	alternativeValSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	last := len(alternativeValSetChanges) - 1
	alternativeValSetChanges[last].Validators = valList.Validators[6:8]
	alternativeValSetChanges[last].PrivVals = valList.PrivVals[6:8]
	testBisection, _, _, _, _ := generateMultiPeerBisectionCase(
		description,
		primaryValSetChanges,
		alternativeValSetChanges,
		2,
		expectedOutputNoError,
	)

	testBisection.Witnesses = append(testBisection.Witnesses, testBisection.Primary)

	file := MULTI_PEER_BISECTION_PATH + "conflicting_valid_commits_from_one_of_the_witnesses.json"
	testBisection.genJSON(file)
}

// Also a case where some validators have double signed
func caseBisectionConflictingHeaders(valList ValList) {
	description := "Case: Trusted height=1, bisecting to verify height=5 and receives a conflicting header from witness, should expect error"
	copiedValList := valList.Copy()
	valSetChanges := ValSetChanges{}.getDefault(copiedValList)[:4]
	lastValSetChange := ValList{
		copiedValList.Validators[:4],
		copiedValList.PrivVals[:4],
	}
	valSetChanges = append(valSetChanges, lastValSetChange)

	testBisection, _, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		2,
	)
	last := len(testBisection.Primary.LiteBlocks) - 1
	testBisection.Primary.LiteBlocks[last].SignedHeader.Commit.Signatures[0] = types.CommitSig{
		BlockIDFlag:      types.BlockIDFlagAbsent,
		ValidatorAddress: nil,
	}

	testBisection2, states, privVals := generateGeneralBisectionCase(
		description,
		valSetChanges,
		2,
	)

	state := states[len(states)-2]
	state.Validators.IncrementProposerPriority(1)
	lastCommit := testBisection2.Primary.LiteBlocks[last-1].SignedHeader.Commit
	time := testBisection2.Primary.LiteBlocks[last-1].SignedHeader.Header.Time.Add(2 * time.Second)
	liteBlock, _, _ := generateNextBlock(state, privVals, lastCommit, time)
	liteBlock.SignedHeader.Commit.Signatures[1] = types.CommitSig{
		BlockIDFlag:      types.BlockIDFlagAbsent,
		ValidatorAddress: nil,
	}

	testBisection2.Primary.LiteBlocks[last] = liteBlock
	testBisection.Witnesses[0] = testBisection2.Primary

	testBisection.HeightToVerify = 5
	testBisection.ExpectedOutput = expectedOutputError

	file := MULTI_PEER_BISECTION_PATH + "conflicting_headers.json"
	testBisection.genJSON(file)
}
