package generator

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// This is the information we get from the TLC output
// It is the minimum details and the rest needs to be mocked out before generating a JSON test case from this
type tlcOutput struct {
	Height           int
	MinTrustedHeight int
	Blockchain       []block
	Verdict          bool
}

// This represents what a block looks like in a TLC output
type block struct {
	Height         int
	Validators     validators
	NextValidators validators
	LastCommit     validators
}

// A validator from the TLC output
type validator struct {
	ID          int
	VotingPower int64
}

type validators []validator

func getTLCOutput(file string) tlcOutput {

	// reads lines from the text file
	lines, err := scanLines(file)
	if err != nil {
		panic(err)
	}

	var tlcOutput tlcOutput
	pos1 := "$1"

	// iterates through each line and looks for the necessary data that we need to fill out the TLCOutput struct
	for _, line := range lines {
		// looks for the height we're at
		heightSearch := regexp.MustCompile(`/\\ height = (\d)`)
		height, found := searchAndExtract(heightSearch, line, pos1)
		if found == true {
			tlcOutput.Height = stringToInt(height)
		}

		// looks for the minimum trusted height
		minTrustedHeightSearch := regexp.MustCompile(`/\\ minTrustedHeight = (\d)`)
		minTrustedHeight, found := searchAndExtract(minTrustedHeightSearch, line, pos1)
		if found == true {
			tlcOutput.MinTrustedHeight = stringToInt(minTrustedHeight)
		}

		// looks for the verdict it ended with
		verdictSearch := regexp.MustCompile(`\bverdict \|-> (\w{4,5})`)
		verdict, found := searchAndExtract(verdictSearch, line, pos1)
		if found == true {
			tlcOutput.Verdict = stringToBool(verdict)
		}

		// looks for "blockchain" which consists of a list of stored
		// headers, validator sets and commits from each height
		blockchainSearch := regexp.MustCompile(`/\\ blockchain = .*`)
		blockchain, found := search(blockchainSearch, line)
		if found == true {

			headerHeightSearch := regexp.MustCompile(`height \|-> (\d)`)
			heights := searchAllAndExtractAll(headerHeightSearch, blockchain, pos1)

			vpIDs, VPs := nestedSearchAndDoubleExtract(`\WVP \|-> \(A_p\d :> \d( @{2} A_p\d :> \d){0,4}\)`, `A_p(\d) :> (\d)`, blockchain, pos1, "$2")

			nextvpIDs, nextVPs := nestedSearchAndDoubleExtract(`NextVP \|-> \(A_p\d :> \d( @{2} A_p\d :> \d){0,4}\)`, `A_p(\d) :> (\d)`, blockchain, pos1, "$2")

			lastCommits := nestedSearchAndExtract(`lastCommit \|-> \{(A_p\d)?(, A_p\d){0,4}\}`, `A_p(\d)`, blockchain, pos1)

			tlcOutput.Blockchain = makeBlockchain(heights, vpIDs, VPs, nextvpIDs, nextVPs, lastCommits)

		}
	}
	// print(tlcOutput)

	return tlcOutput
}

// prints the given TLCOutput in a more readable format
func print(tlcOutput tlcOutput) {
	fmt.Println("Height: ", tlcOutput.Height)
	fmt.Println("MinTrustedHeight: ", tlcOutput.MinTrustedHeight)
	fmt.Println("Blockchain: ")
	for _, block := range tlcOutput.Blockchain {
		fmt.Printf("\n%+v\n", block)
	}
	fmt.Println("\nVerdict: ", tlcOutput.Verdict)
}

// Searches for the 'searchLine' in 'line'
// and calls SearchAllAndExtractAll function to search all the matching 'searchItem'
// and extract the value at given position 'pos'
func nestedSearchAndExtract(searchLine string, searchItem string, line string, pos string) (result [][]string) {
	searchLn := regexp.MustCompile(searchLine)
	searchItm := regexp.MustCompile(searchItem)
	items := searchAll(searchLn, line)
	for _, itm := range items {
		result = append(result, searchAllAndExtractAll(searchItm, itm, pos))
	}
	return
}

// similar to NestedSearchAndExtract
// but here you can extract two values from the 'searchItem'
// by specifying the two positions, namely, 'pos1' and 'pos2'
func nestedSearchAndDoubleExtract(searchLine string, searchItem string, line string, pos1 string, pos2 string) (result1, result2 [][]string) {
	searchLn := regexp.MustCompile(searchLine)
	searchItm := regexp.MustCompile(searchItem)
	items := searchAll(searchLn, line)
	for _, itm := range items {
		result1 = append(result1, searchAllAndExtractAll(searchItm, itm, pos1))
		result2 = append(result2, searchAllAndExtractAll(searchItm, itm, pos2))

	}
	return
}

// returns a list of block that represents a blockchain
// it takes in a list og heights, validator ids, validators, next validator ids, next validators and last commits
// this function makes a block identifying data at index 0 to be data for first block
func makeBlockchain(heights []string, vpIDs [][]string, VPs [][]string, nextvpIDs [][]string, nextVPs [][]string, lastCommits [][]string) (blockchain []block) {

	heightInt := sliceStringToInt(heights)

	for i, height := range heightInt {
		var block block
		block.Height = height

		for j, eachID := range vpIDs[i] {
			VPsIntSlice := sliceStringToInt(VPs[i])
			// vpIDsIntSlice := SliceStringToInt(vpIDs[i])
			var validator validator
			validator.ID = stringToInt(eachID)
			validator.VotingPower = int64(VPsIntSlice[j])
			block.Validators = append(block.Validators, validator)
		}

		for j, eachID := range nextvpIDs[i] {
			nextVPsIntSlice := sliceStringToInt(nextVPs[i])
			var nextValidator validator
			nextValidator.ID = stringToInt(eachID)
			nextValidator.VotingPower = int64(nextVPsIntSlice[j])
			block.NextValidators = append(block.NextValidators, nextValidator)
		}

		for _, eachID := range lastCommits[i] {
			var lastCommit validator
			prevVals := blockchain[i-1].Validators
			lastCommit = prevVals.getValidatorByID(stringToInt(eachID))
			// fmt.Println(lastCommit)
			block.LastCommit = append(block.LastCommit, lastCommit)
		}
		blockchain = append(blockchain, block)
	}
	return
}

// returns the validator that has the given id
func (vals validators) getValidatorByID(id int) (reqVal validator) {
	for _, val := range vals {
		if val.ID == id {
			reqVal = val
		}
	}
	return
}

// converts a slice of string to a slice of integer
// this is needed because numbers are read in string from the tlc output
func sliceStringToInt(strSlice []string) (intSlice []int) {
	for _, str := range strSlice {
		i := stringToInt(str)
		intSlice = append(intSlice, i)
	}
	return
}

// converts string to bool
// i.e. if the string is "TRUE" it returns a bool value of true
func stringToBool(str string) bool {
	if str == "TRUE" {
		return true
	}
	return false
}

// converts string to integer
func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

// searches for ALL that matches 'search' in the line 'line'
// and extracts ALL the value at position 'pos' from the result of the search
func searchAllAndExtractAll(search *regexp.Regexp, line string, pos string) (result []string) {
	match := searchAll(search, line)
	result = extractAll(search, pos, match)
	return
}

// searches for a match of 'search' in the line 'line'
// and extracts value at position 'pos'
func searchAndExtract(srch *regexp.Regexp, line string, pos string) (string, bool) {
	match, found := search(srch, line)
	if found == true {
		result := extract(srch, pos, match)
		return result, true
	}
	return "", false
}

// searches for a match of 'search' in the line 'line'
// returns the match, if found, with true
// otherwise returns empty string with false
func search(search *regexp.Regexp, line string) (string, bool) {
	match := search.FindString(line)
	if match != "" {
		return match, true
	}
	return "", false
}

// similar to Search function
// but here it searches ALL the matches, not just the first found match
// returns a slice of matches found
func searchAll(search *regexp.Regexp, line string) []string {
	match := search.FindAllString(line, -1)
	return match
}

// extracts the value at position 'pos' in the string 'match' which is within 'search' regexp
func extract(search *regexp.Regexp, pos string, match string) string {
	result := search.ReplaceAllString(match, pos)
	return result
}

// similar to Extract function
// but here it extracts values from a list of string 'match'
func extractAll(search *regexp.Regexp, pos string, match []string) (result []string) {
	for _, m := range match {
		result = append(result, search.ReplaceAllString(m, pos))
	}
	return
}

// scans lines from the text file specified by the 'path'
func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
