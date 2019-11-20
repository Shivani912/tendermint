package generator

import (
	"time"

	"github.com/tendermint/tendermint/version"
)

func CaseVerifyHeaderEmpty(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}
	name := "verify"
	description := "Case: one lite block, empty header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header = nil
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongLastCommitHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong last commit hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.LastCommitHash = []byte("wrong hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongLastResultsHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong last results hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.LastResultsHash = []byte("wrong hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}
func CaseVerifyHeaderWrongLastBlockID(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong last block ID in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.LastBlockID.Hash = []byte("wrong hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongDataHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong data hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.DataHash = []byte("wrong hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongChainID(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong chain ID in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.ChainID = "wrong chain id"
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongVersion(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong version in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.Version.Block = version.Protocol(1)
	testCase.Input[0].SignedHeader.Header.Version.App = version.Protocol(1)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongHeight(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong height in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.Height += 1
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongTimestamp(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong timestamp in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.Time = secondBlockTime.Add(1 * time.Minute)
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongNumTxs(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong number of transactions in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.NumTxs += 1
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongTotalTxs(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong number of total transactions in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.TotalTxs += 1
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongValSetHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong val set hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.ValidatorsHash = []byte("wrong validator set hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongNextValSetHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong next val set hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.NextValidatorsHash = []byte("wrong next validator set hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}

func CaseVerifyHeaderWrongConsensusHash(testCases *TestCases, valList ValList) {
	var testCase *TestCase = &TestCase{}

	name := "verify"
	description := "Case: one lite block, wrong consensus hash in header, with error"

	GenerateGeneralTestCase(testCase, valList, 1, name, description)
	testCase.Input[0].SignedHeader.Header.ConsensusHash = []byte("wrong consensus hash")
	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}
