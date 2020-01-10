package generator

import "time"

func caseVerifyHeaderEmpty(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, empty header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header = nil
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongLastCommitHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last commit hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastCommitHash = []byte("wrong hash")
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongLastResultsHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last results hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastResultsHash = []byte("wrong hash")
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
func caseVerifyHeaderWrongLastBlockID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong last block ID in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastBlockID.Hash = []byte("wrong hash")
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongChainID(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong chain ID in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ChainID = "wrong chain id"
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongHeight(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong height in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Height++
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongTimestamp(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong timestamp in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Time = secondBlockTime.Add(1 * time.Minute)
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongValSetHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong val set hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ValidatorsHash = []byte("wrong validator set hash")
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}

func caseVerifyHeaderWrongNextValSetHash(testBatch *TestBatch, valList ValList) {

	description := "Case: one lite block, wrong next val set hash in header, with error"

	initial, input, _, _ := generateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.NextValidatorsHash = []byte("wrong next validator set hash")
	testCase := generateTestCase(description, initial, input, expectedOutputError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
