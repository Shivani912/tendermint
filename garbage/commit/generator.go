package main

import (

	// "fmt"
	// "encoding/json"
	// "io/ioutil"

	"github.com/tendermint/tendermint/garbage"
)

func main() {

	garbage.GenerateTestCase()
	// commit := types.RandCommit(5,1,4,3)

	// j, err := json.MarshalIndent(commit, "", "	")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var c types.Commit

	// json.Unmarshal(j, &c)

	// file := "./commit.json"
	// _ = ioutil.WriteFile(file, j, 0644)
	// fmt.Println(c)
}
