# ADR 046: Language Agnostic Testing

## Changelog
* 07-11-2019: Initial Draft

## Context

Language Agnostic Test Suites are a great way to outline the expected behaviour of a system in a given situation. As we are working towards having multiple implementations of Tendermint (Go, Rust, Scala and maybe more in future), having a general english spec is obviously not enough. A structured JSON specification as a base for testing can help greatly in ensuring the functional behaviour of Tendermint to be same across different implementations.

## Decision

Considering Go implementation to be the standard, Go code is used to generate JSON test suites. Generation is sort of locally simulating the process (that needs to be tested) and capturing the data involved (input/output) into JSON format. 

A generator will look something like this: 
```
func CaseVerifyValidatorSetOf1(testCases *TestCases, valList *ValList) {
	var testCase *TestCase = &TestCase{}
	GenerateTestNameAndDescription(testCase, "verify", "Case: one lite block, one validator, no error")

	vals := valList.ValidatorSet.Validators[:1]
	privVal := valList.PrivVal[:1]

	state := GenerateFirstBlock(testCase, vals, privVal, now)
	GenerateNextBlock(state, testCase, privVal, testCase.Initial.SignedHeader.Commit, now.Add(time.Second*5))
	GenerateInitial(testCase, testCase.Input[0].ValidatorSet, 3*time.Hour, now.Add(time.Second*10))

	GenerateExpectedOutput(testCase)
	testCases.TC = append(testCases.TC, *testCase)
}
```

Each test case contains a description, what it is testing, some initial state to begin with, input data and the expected output.

```
{
    "name": "",
    "description": "",
    "initial": {},
    "input": {},
    "expected_output": ""
}
```

The tests must be able to identify the test name and expect that the initial and input data will result into the expected output.

## Status

Accepted

## Consequences


### Positive

### Negative

### Neutral

This will not require any change in the main code base and therefore, does not have any direct consequences. It is useful mainly for testing the Rust implementation.

