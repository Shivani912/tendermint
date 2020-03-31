package generator

func caseBisectionConflictingValidCommits(valList ValList) {
	description := "Case: Trusted height=1, bisecting to verify height=11, should not expect error"
	valSetChanges := ValSetChanges{}.getDefault(valList.Copy())
	testBisection, _, _ := generateGeneralBisectionCase(
		description,
		valSetChanges,
		int32(2),
	)

	file := "./tests/json/multi_peer_bisection/conflicting_valid_commits.json"
	testBisection.genJSON(file)

}
