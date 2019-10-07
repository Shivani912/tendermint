package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"io/ioutil"

	amino "github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"github.com/tendermint/tendermint/types"
)


func main() {
	mode := os.Args[1]
	folder := os.Args[2]	
	
	switch mode {
	case "validator":

		num := os.Args[3]

		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		
		GenerateValidator(folder, numInt)
		fmt.Println("Validator files generated!")

	case "merkle":
		calculateMerkleRoot(folder)
		fmt.Println("merkle root file generated!")

	default:
		fmt.Println("invalid command")
	}

}

func GenerateValidator(folder string, numInt int) {
	for i := 1; i <= numInt; i++ {

		v, _ := types.RandValidator(true , 1)

		var cdc = amino.NewCodec()
		cryptoAmino.RegisterAmino(cdc)

		b, err := cdc.MarshalJSON(v)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		num := strconv.Itoa(i)
		file := folder+"/val"+num+".json"
	
		_ = ioutil.WriteFile(file, b, 0644)
	}
}

func calculateMerkleRoot(folder string) {

	valSet := getValidatorSet(folder)

	var cdc = amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)

	mr := string(toHex(valSet.Hash()))

	var merkle Merkle
	merkle.MerkleRoot = mr

	fmt.Println(merkle.MerkleRoot)

	// if mode == "file" {
		b, err := json.Marshal(merkle)
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		file := folder+"/merkle_root.json"

		_ = ioutil.WriteFile(file, b, 0644)
		// string(toHex(valSet.Hash()))

	// }else if mode == "bytes" {
	// 	return
	// }

	

}
