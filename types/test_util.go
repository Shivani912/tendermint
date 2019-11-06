package types

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/tmhash"

	cmn "github.com/tendermint/tendermint/libs/common"
)

// var now, _ = time.Parse(time.RFC3339, "2019-11-02T15:04:05Z")

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
			Timestamp:        now, //tmtime.Now(),
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

func MakeVote(height int64, blockID BlockID, valSet *ValidatorSet, privVal PrivValidator, chainID string, now time.Time) (*Vote, error) {
	addr := privVal.GetPubKey().Address()
	idx, _ := valSet.GetByAddress(addr)
	vote := &Vote{
		ValidatorAddress: addr,
		ValidatorIndex:   idx,
		Height:           height,
		Round:            0,
		Timestamp:        now, //tmtime.Now(),
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
			NumTxs: int64(len(txs)),
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

// ----- Vote_set test utils -----

// NOTE: privValidators are in order
func randVoteSet(height int64, round int, type_ SignedMsgType, numValidators int, votingPower int64) (*VoteSet, *ValidatorSet, []PrivValidator) {
	valSet, privValidators := RandValidatorSet(numValidators, votingPower)
	return NewVoteSet("test_chain_id", height, round, type_, valSet), valSet, privValidators
}

// Convenience: Return new vote with different validator address/index
func withValidator(vote *Vote, addr []byte, idx int) *Vote {
	vote = vote.Copy()
	vote.ValidatorAddress = addr
	vote.ValidatorIndex = idx
	return vote
}

// Convenience: Return new vote with different height
func withHeight(vote *Vote, height int64) *Vote {
	vote = vote.Copy()
	vote.Height = height
	return vote
}

// Convenience: Return new vote with different round
func withRound(vote *Vote, round int) *Vote {
	vote = vote.Copy()
	vote.Round = round
	return vote
}

// Convenience: Return new vote with different type
func withType(vote *Vote, type_ byte) *Vote {
	vote = vote.Copy()
	vote.Type = SignedMsgType(type_)
	return vote
}

// Convenience: Return new vote with different blockHash
func withBlockHash(vote *Vote, blockHash []byte) *Vote {
	vote = vote.Copy()
	vote.BlockID.Hash = blockHash
	return vote
}

// Convenience: Return new vote with different blockParts
func withBlockPartsHeader(vote *Vote, blockPartsHeader PartSetHeader) *Vote {
	vote = vote.Copy()
	vote.BlockID.PartsHeader = blockPartsHeader
	return vote
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

// ----- Validator set test utils -----

func newValidator(address []byte, power int64) *Validator {
	return &Validator{Address: address, VotingPower: power}
}

func randPubKey() crypto.PubKey {
	var pubKey [32]byte
	copy(pubKey[:], cmn.RandBytes(32))
	return ed25519.PubKeyEd25519(pubKey)
}

func randValidator_(totalVotingPower int64) *Validator {
	// this modulo limits the ProposerPriority/VotingPower to stay in the
	// bounds of MaxTotalVotingPower minus the already existing voting power:
	val := NewValidator(randPubKey(), int64(cmn.RandUint64()%uint64((MaxTotalVotingPower-totalVotingPower))))
	val.ProposerPriority = cmn.RandInt64() % (MaxTotalVotingPower - totalVotingPower)
	return val
}

func randValidatorSet(numValidators int) *ValidatorSet {
	validators := make([]*Validator, numValidators)
	totalVotingPower := int64(0)
	for i := 0; i < numValidators; i++ {
		validators[i] = randValidator_(totalVotingPower)
		totalVotingPower += validators[i].VotingPower
	}
	return NewValidatorSet(validators)
}

func (valSet *ValidatorSet) toBytes() []byte {
	bz, err := cdc.MarshalBinaryLengthPrefixed(valSet)
	if err != nil {
		panic(err)
	}
	return bz
}

func (valSet *ValidatorSet) fromBytes(b []byte) {
	err := cdc.UnmarshalBinaryLengthPrefixed(b, &valSet)
	if err != nil {
		// DATA HAS BEEN CORRUPTED OR THE SPEC HAS CHANGED
		panic(err)
	}
}

type testVal struct {
	name  string
	power int64
}

func permutation(valList []testVal) []testVal {
	if len(valList) == 0 {
		return nil
	}
	permList := make([]testVal, len(valList))
	perm := cmn.RandPerm(len(valList))
	for i, v := range perm {
		permList[v] = valList[i]
	}
	return permList
}

func createNewValidatorList(testValList []testVal) []*Validator {
	valList := make([]*Validator, 0, len(testValList))
	for _, val := range testValList {
		valList = append(valList, newValidator([]byte(val.name), val.power))
	}
	return valList
}

func createNewValidatorSet(testValList []testVal) *ValidatorSet {
	return NewValidatorSet(createNewValidatorList(testValList))
}

func valSetTotalProposerPriority(valSet *ValidatorSet) int64 {
	sum := int64(0)
	for _, val := range valSet.Validators {
		// mind overflow
		sum = safeAddClip(sum, val.ProposerPriority)
	}
	return sum
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

// ----- tx test utils -----

func makeTxs(cnt, size int) Txs {
	txs := make(Txs, cnt)
	for i := 0; i < cnt; i++ {
		txs[i] = cmn.RandBytes(size)
	}
	return txs
}

func randInt(low, high int) int {
	off := cmn.RandInt() % (high - low)
	return low + off
}

// ----- block test utils -----

func makeBlockIDRandom() BlockID {
	blockHash := make([]byte, tmhash.Size)
	partSetHash := make([]byte, tmhash.Size)
	rand.Read(blockHash)   //nolint: gosec
	rand.Read(partSetHash) //nolint: gosec
	blockPartsHeader := PartSetHeader{123, partSetHash}
	return BlockID{blockHash, blockPartsHeader}
}

func makeBlockID(hash []byte, partSetSize int, partSetHash []byte) BlockID {
	return BlockID{
		Hash: hash,
		PartsHeader: PartSetHeader{
			Total: partSetSize,
			Hash:  partSetHash,
		},
	}

}

func randCommit(now time.Time) *Commit {
	lastID := makeBlockIDRandom()
	h := int64(3)
	voteSet, _, vals := randVoteSet(h-1, 1, PrecommitType, 10, 1)
	commit, err := MakeCommit(lastID, h-1, 1, voteSet, vals, now)
	if err != nil {
		panic(err)
	}
	return commit
}

// Exposed functions for testing purpose

func ExamplePrevote() *Vote {
	return examplePrevote()
}

func ExamplePrecommit() *Vote {
	return examplePrecommit()
}

// ---- generator utils ----

func toHex(mr []byte) []byte {
	src := mr

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}
