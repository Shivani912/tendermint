package generator

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

func caseBisectionHappyPath(valList ValList) {
	description := "Case: Trusted height=1, bisecting to verify height=11, should not expect error"
	valSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	testBisection, _, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		int32(2),
	)

	file := "./tests/json/many_header_bisection/happy_path.json"
	testBisection.genJSON(file)

}

func caseBisectionWorstCase(valList ValList) {
	description := "Case: Trusted height=1, bisecting at every height, should not expect error"

	copiedValList := valList.Copy()
	valsArray := [][]*types.Validator{
		copiedValList.Validators[0:1],
		copiedValList.Validators[1:2],
		copiedValList.Validators[2:3],
		copiedValList.Validators[3:4],
		copiedValList.Validators[4:5],
		copiedValList.Validators[5:6],
		copiedValList.Validators[6:7],
		copiedValList.Validators[7:8],
		copiedValList.Validators[8:9],
		copiedValList.Validators[9:10],
		copiedValList.Validators[10:11],
	}
	privValsArray := []types.PrivValidatorsByAddress{
		copiedValList.PrivVals[0:1],
		copiedValList.PrivVals[1:2],
		copiedValList.PrivVals[2:3],
		copiedValList.PrivVals[3:4],
		copiedValList.PrivVals[4:5],
		copiedValList.PrivVals[5:6],
		copiedValList.PrivVals[6:7],
		copiedValList.PrivVals[7:8],
		copiedValList.PrivVals[8:9],
		copiedValList.PrivVals[9:10],
		copiedValList.PrivVals[10:11],
	}

	valSetChanges := ValSetChanges{}.makeValSetChanges(valsArray, privValsArray)

	testBisection, _, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		10,
	)

	file := "./tests/json/many_header_bisection/worst_case.json"
	testBisection.genJSON(file)
}

func caseBisectionInvalidValidatorSet(valList ValList) {
	description := "Case: Trusted height = 1, fails at height 6 on finding invalid validator set"

	valSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	testBisection, states, privVals := generateGeneralBisectionCase(
		description,
		valSetChanges,
		2,
	)
	valSet := types.NewValidatorSet(valList.Validators[15:16])
	lastCommit := testBisection.Primary.LiteBlocks[5].SignedHeader.Commit
	time := testBisection.Primary.LiteBlocks[5].SignedHeader.Commit.Signatures[0].Timestamp

	state := states[len(states)-1]
	state.Validators = valSet
	privVals = valList.PrivVals[15:16]
	liteBlock, _, _ := generateNextBlock(state, privVals, lastCommit, time)

	testBisection.Primary.LiteBlocks[5] = liteBlock
	testBisection.ExpectedOutput = expectedOutputError

	file := "./tests/json/many_header_bisection/invalid_validator_set.json"
	testBisection.genJSON(file)
}

func caseBisectionNotEnoughCommits(valList ValList) {
	description := "Case: Trusted height=1, fails at height 6 because more than one-third (trust level) vals didn't sign"

	copiedValList := valList.Copy()

	valsArray := [][]*types.Validator{
		valList.Validators[:4],
		valList.Validators[:4],
		valList.Validators[:4],
		valList.Validators[:4],
		valList.Validators[:4],
	}

	privValsArray := []types.PrivValidatorsByAddress{
		valList.PrivVals[:4],
		valList.PrivVals[:4],
		valList.PrivVals[:4],
		valList.PrivVals[:4],
		valList.PrivVals[:4],
	}

	valSetChanges := ValSetChanges{}.makeValSetChanges(valsArray, privValsArray)
	testBisection, states, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		3,
	)

	valsArray = [][]*types.Validator{
		copiedValList.Validators[0:1],
		copiedValList.Validators[0:1],
		copiedValList.Validators[0:1],
	}

	privValsArray = []types.PrivValidatorsByAddress{
		copiedValList.PrivVals[0:1],
		copiedValList.PrivVals[0:1],
		copiedValList.PrivVals[0:1],
	}

	valSetChanges = ValSetChanges{}.makeValSetChanges(valsArray, privValsArray)
	last := len(testBisection.Primary.LiteBlocks) - 1
	lastCommit := testBisection.Primary.LiteBlocks[last].SignedHeader.Commit
	blockTime := lastCommit.Signatures[0].Timestamp.Add(5 * time.Second)
	last = len(states) - 1
	state := states[last]
	state.Validators = types.NewValidatorSet(copiedValList.Validators[0:1])
	privVals := copiedValList.PrivVals[0:1]
	liteBlocks, _, _ := generateNextBlocks(3, state, privVals, lastCommit, valSetChanges, blockTime)
	testBisection.Primary.LiteBlocks = append(testBisection.Primary.LiteBlocks, liteBlocks...)
	testBisection.Witnesses[0] = testBisection.Primary
	testBisection.HeightToVerify = 8
	testBisection.ExpectedOutput = expectedOutputError

	file := "./tests/json/many_header_bisection/not_enough_commits.json"
	testBisection.genJSON(file)

}

func caseBisectionHeaderOutOfTrustingPeriod(valList ValList) {
	description := "Case: Trusted height=1, fails at height 11 because header at height 1 runs out of trusting period while bisecting"
	valSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	testBisection, _, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		int32(1),
	)

	testBisection.TrustOptions.Period = 30 * time.Second
	testBisection.ExpectedOutput = expectedOutputError

	file := "./tests/json/many_header_bisection/header_out_of_trusting_period.json"
	testBisection.genJSON(file)
}
