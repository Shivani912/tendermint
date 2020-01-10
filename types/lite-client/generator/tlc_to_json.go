package generator

import (
	"fmt"

	"github.com/tendermint/tendermint/types"
)

// should return TestCase
func tlcOutputToTestCase(tlcOutput tlcOutput, valList ValList) {

	findInitial(tlcOutput.MinTrustedHeight, tlcOutput.Blockchain, valList)
}

func findInitial(minTrustedHeight int, blockchain []block, valList ValList) (initial Initial) {

	for _, block := range blockchain {
		if block.Height == minTrustedHeight {

			if block.LastCommit == nil {
				fmt.Println("nil")
			}
			valSet := makeValSet(block.LastCommit, valList)
			fmt.Println(block.LastCommit, valSet)
			// initial = GenerateInitialAtHeight(minTrustedHeight, valSet)
		}
	}
	return
}

func makeValSet(valz validators, valList ValList) types.ValidatorSet {

	copyValList := valList.Copy()

	var valzFromValList []*types.Validator
	for _, val := range valz {
		valToBeAdded := copyValList.Validators[val.ID]
		valToBeAdded.VotingPower = val.VotingPower
		valzFromValList = append(valzFromValList)
	}
	valSet := *types.NewValidatorSet(valzFromValList)
	return valSet
}
