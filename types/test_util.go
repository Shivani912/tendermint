package types

import (
	"time"
)

func MakeCommit(blockID BlockID, height int64, round int,
	voteSet *VoteSet, validators []PrivValidator, now time.Time) (*Commit, error) {

	// all sign
	for i := 0; i < len(validators); i++ {
		addr := validators[i].GetPubKey().Address()
		vote := &Vote{
			ValidatorAddress: addr,
			ValidatorIndex:   i,
			Height:           height,
			Round:            round,
			Type:             PrecommitType,
			BlockID:          blockID,
			Timestamp:        now,
		}

		_, err := signAddVote(validators[i], vote, voteSet)
		if err != nil {
			return nil, err
		}
	}

	return voteSet.MakeCommit(), nil
}

func signAddVote(privVal PrivValidator, vote *Vote, voteSet *VoteSet) (signed bool, err error) {
	err = privVal.SignVote(voteSet.ChainID(), vote)
	if err != nil {
		return false, err
	}
	return voteSet.AddVote(vote)
}

func MakeVote(
	height int64,
	blockID BlockID,
	valSet *ValidatorSet,
	privVal PrivValidator,
	chainID string,
	now time.Time,
) (*Vote, error) {
	addr := privVal.GetPubKey().Address()
	idx, _ := valSet.GetByAddress(addr)
	vote := &Vote{
		ValidatorAddress: addr,
		ValidatorIndex:   idx,
		Height:           height,
		Round:            0,
		Timestamp:        now,
		Type:             PrecommitType,
		BlockID:          blockID,
	}
	if err := privVal.SignVote(chainID, vote); err != nil {
		return nil, err
	}
	return vote, nil
}

// MakeBlock returns a new block with an empty header, except what can be
// computed from itself.
// It populates the same set of fields validated by ValidateBasic.
func MakeBlock(height int64, txs []Tx, lastCommit *Commit, evidence []Evidence) *Block {
	block := &Block{
		Header: Header{
			Height: height,
		},
		Data: Data{
			Txs: txs,
		},
		Evidence:   EvidenceData{Evidence: evidence},
		LastCommit: lastCommit,
	}
	block.fillHeader()
	return block
}

// ----- Vote test utils ------

func examplePrevote() *Vote {
	return exampleVote(byte(PrevoteType))
}

func examplePrecommit() *Vote {
	return exampleVote(byte(PrecommitType))
}

func exampleVote(t byte) *Vote {
	var stamp, err = time.Parse(TimeFormat, "2017-12-25T03:00:01.234Z")
	if err != nil {
		panic(err)
	}

	return &Vote{
		Type:      SignedMsgType(t),
		Height:    12345,
		Round:     2,
		Timestamp: stamp,
		BlockID: BlockID{
			Hash: tmhash.Sum([]byte("blockID_hash")),
			PartsHeader: PartSetHeader{
				Total: 1000000,
				Hash:  tmhash.Sum([]byte("blockID_part_set_header_hash")),
			},
		},
		ValidatorAddress: crypto.AddressHash([]byte("validator_address")),
		ValidatorIndex:   56789,
	}
}


func toTestValList(valList []*Validator) []testVal {
	testList := make([]testVal, len(valList))
	for i, val := range valList {
		testList[i].name = string(val.Address)
		testList[i].power = val.VotingPower
	}
	return testList
}

func testValSet(nVals int, power int64) []testVal {
	vals := make([]testVal, nVals)
	for i := 0; i < nVals; i++ {
		vals[i] = testVal{fmt.Sprintf("v%d", i+1), power}
	}
	return vals
}

type valSetErrTestCase struct {
	startVals  []testVal
	updateVals []testVal
}

// Sort validators by priority and address
type validatorsByPriority []*Validator

func (valz validatorsByPriority) Len() int {
	return len(valz)
}

func (valz validatorsByPriority) Less(i, j int) bool {
	if valz[i].ProposerPriority < valz[j].ProposerPriority {
		return true
	}
	if valz[i].ProposerPriority > valz[j].ProposerPriority {
		return false
	}
	return bytes.Compare(valz[i].Address, valz[j].Address) < 0
}

func (valz validatorsByPriority) Swap(i, j int) {
	it := valz[i]
	valz[i] = valz[j]
	valz[j] = it
}

//-------------------------------------
// Sort testVal-s by address.
type testValsByAddress []testVal

func (tvals testValsByAddress) Len() int {
	return len(tvals)
}

func (tvals testValsByAddress) Less(i, j int) bool {
	return bytes.Compare([]byte(tvals[i].name), []byte(tvals[j].name)) == -1
}

func (tvals testValsByAddress) Swap(i, j int) {
	it := tvals[i]
	tvals[i] = tvals[j]
	tvals[j] = it
}