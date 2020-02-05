package generator

func caseBisectionVerifyTenHeaders(testBatch *TestBatch, valList ValList) {
	description := "Case: Trusted height=1, bisecting to verify height=11, should not expect error"

	initial, input, state, privVals := generateGeneralCase(valList, 3)
	lastCommit := input[0].SignedHeader.Commit
	liteBlocks, _, _ := generateNextBlocks(10, state, privVals, lastCommit)
	input = append(input, liteBlocks...)
	testCase := makeTestCase(description, initial, input, expectedOutputNoError)
	testBatch.TestCases = append(testBatch.TestCases, testCase)
}
