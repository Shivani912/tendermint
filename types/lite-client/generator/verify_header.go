package generator

import "time"

func CaseVerifyHeaderEmpty(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, empty header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header = nil
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongLastCommitHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong last commit hash in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastCommitHash = []byte("wrong hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongLastResultsHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong last results hash in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastResultsHash = []byte("wrong hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}
func CaseVerifyHeaderWrongLastBlockID(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong last block ID in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.LastBlockID.Hash = []byte("wrong hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongChainID(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong chain ID in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ChainID = "wrong chain id"
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongHeight(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong height in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Height += 1
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongTimestamp(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong timestamp in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.Time = secondBlockTime.Add(1 * time.Minute)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongValSetHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong val set hash in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.ValidatorsHash = []byte("wrong validator set hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}

func CaseVerifyHeaderWrongNextValSetHash(testCases *TestCases, valList ValList) {

	description := "Case: one lite block, wrong next val set hash in header, with error"

	initial, input, _, _ := GenerateGeneralCase(valList, 1)
	input[0].SignedHeader.Header.NextValidatorsHash = []byte("wrong next validator set hash")
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputError)
	testCases.TC = append(testCases.TC, testCase)
}
