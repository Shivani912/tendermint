package generator

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
