package types

import (
	"fmt"
	"time"
)

func GenerateCommit(header Header, partSet *PartSet, valSet ValidatorSet, privVal []PrivValidator, chainID string, now time.Time) *Commit {
	blockID := &BlockID{
		Hash: header.Hash(),
		PartsHeader: PartSetHeader{
			Hash:  partSet.Hash(),
			Total: partSet.Total(),
		},
	}
	voteSet := NewVoteSet(chainID, header.Height, 1, SignedMsgType(byte(PrecommitType)), &valSet)

	commit, err := MakeCommit(*blockID, header.Height, 1, voteSet, privVal, now)
	if err != nil {
		fmt.Println(err)
	}
	return commit
}

func GenerateTxs() []Tx {
	// Empty txs
	return []Tx{}
}

func GenerateEvidences() []Evidence {
	// Empty evidences
	return []Evidence{}
}
