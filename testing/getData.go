package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	amino "github.com/tendermint/go-amino"

	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
	types "github.com/tendermint/tendermint/types"
)

type Merkle struct {
	MerkleRoot      string `json:"merkle_root"`
	NumOfValidators int    `json:"num_of_validators"`
}

func getData(folder string) (*types.ValidatorSet, string) {

	var merkle Merkle
	file := fmt.Sprintf("%smerkle_root.json", folder)
	merklleJson := getJsonFrom(file)
	json.Unmarshal(merklleJson, &merkle)
	ValSet := getValidatorSet(merkle.NumOfValidators, folder)

	return ValSet, merkle.MerkleRoot
}
func Equal(a, b []byte) bool {

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func getJsonFrom(file string) []byte {
	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	dat, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return dat
}

func UnmarshalValidator(dat []byte) *types.Validator {
	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)

	var val *types.Validator
	er := cdc.UnmarshalJSON(dat, &val)
	if er != nil {
		fmt.Printf("error: %v", er)
	}
	return val
}

func toHex(mr []byte) []byte {
	src := mr

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func getValidatorSet(num int, folder string) *types.ValidatorSet {

	var file string
	var vals []*types.Validator
	for i := 1; i <= num; i++ {
		file = fmt.Sprintf("%sval%d.json", folder, i)

		dat := getJsonFrom(file)
		val := UnmarshalValidator(dat)

		vals = append(vals, val)
	}

	ValSet := types.NewValidatorSet(vals)
	return ValSet
}
