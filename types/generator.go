package types

import "fmt"


func GenerateCommit(header Header, partSet *PartSet, valSet ValidatorSet, privVal []PrivValidator) *Commit {
	blockID := &BlockID{
		Hash: header.Hash(),
		PartsHeader: PartSetHeader{
			Hash:  partSet.Hash(),
			Total: partSet.Total(),
		},
	}

	voteSet := NewVoteSet("test_chain_id", header.Height, 1, SignedMsgType(byte(PrecommitType)), &valSet)

	commit, err := MakeCommit(*blockID, header.Height, 1, voteSet, privVal)
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
