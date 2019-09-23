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
	// NumOfValidators int    `json:"num_of_validators"`
}

func getData(folder string) (*types.ValidatorSet, string, int) {

	var merkle Merkle
	file := fmt.Sprintf("%s/merkle_root.json", folder)
	merklleJson, err := getJsonFrom(file)
	if err == 1 {
		return nil, "", err
	}
	json.Unmarshal(merklleJson, &merkle)
	ValSet := getValidatorSet( folder)

	return ValSet, merkle.MerkleRoot, 0
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

func getJsonFrom(file string) ([]byte, int) {
	jsonFile, err := os.Open(file)

	if err != nil {

		errStatement := "open "+file+": no such file or directory"
	
		if err.Error() == errStatement {
			return nil, 1
		}else{
			fmt.Println(err)
		}
	}

	defer jsonFile.Close()

	dat, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return dat, 0
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

func getValidatorSet( folder string) *types.ValidatorSet {

	var file string
	var vals []*types.Validator
	for i := 1;; i++ {
		file = fmt.Sprintf("%s/val%d.json", folder, i)

		dat, err := getJsonFrom(file)
		if err == 1 {
			// fmt.Println("error 1")
			break
		}
		val := UnmarshalValidator(dat)

		vals = append(vals, val)
	}

	ValSet := types.NewValidatorSet(vals)
	return ValSet
}
