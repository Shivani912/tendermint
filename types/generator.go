package types

import "fmt"

type TestCase struct {
	LiteBlocks []*LiteBlock `json:"lite_blocks"`
}

type LiteBlock struct {
	SignedHeader *SignedHeader `json:"signed_header"`
	ValidatorSet ValidatorSet  `josn:"validator_set"`
}

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
