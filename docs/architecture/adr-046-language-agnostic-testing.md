# ADR 046: Language Agnostic Testing

## Changelog
* 07-11-2019: Initial Draft
* 02-12-2019: Detailed Explanation

## Context

... TODO: include some background on light client, we're starting there. 
what does a light client need to function? nodes serving headers and commits, and validator sets ...
sequential and bisection. fact of multiple peers.


Currently, we are far from reaching Tendermint v1.0 and expecting quite a few breaking changes on our way. We are positively progressing towards a better version and this means, specs and infrastructure are evolving by every release. While this is the status of Tendermint in Go, we have also begun to work on Tendermint in Rust in parallel, while others are working on implementations in other languages. Moving forward with Rust implementation at this stage can be difficult to keep up with the changes being introduced at every stage in Go implementation.

while we do have eng spec we cant test directly. when we have 2 diff impl we need a way to see it...

To argue this, we do have an English spec that explains the behaviour and implementation of components. However, it clearly does not cover the run time behaviour of data sets. We do follow a test driven development, but again, that does not test integration of multiple components within Tendermint. Also, considering how different coding-style(language specific restrictions and perks, features so we want a common thing that can be used to test that they are compatible) in Rust can be, the implementation will totally differ than in Go. This makes it very easy to develop significant gap between these two implementations.

So to facilitate implementation agnosticism while replicating behaviour of Tendermint, language agnostic tests seem to be the bridge between the gap. It lets you really be precise about how the data should be manipulated and formatted, and decisions to be made within the core logic. You can simulate situations that needs to be tested and define what the expected behaviour should be. There can be multiple levels of testing ranging from unit testing to system-wide testing.

(go impl as reference - we use it to generate but need to be careful of defining the output)
Considering the behaviour outlined by Go implementation to be the standard, we make assertions on that basis but it is not limited to it. Therefore, we also have a good scope on finding existing vulnerabilities(if any) in the system. For producing concrete test cases, we can highly rely on TLC output from Lite Client TLA+ specification. TLC output can be parsed to the JSON test case format, mocking out the remaining data, which makes it acceptable by the main code and can be tested against tests in each implementation.

- we should have files with all the data to test against, eg. using json
- great example is the ethereum test suite ...
- we want to limit the manual work to generate test cases, so ideally we'd be able to use outputs from TLC 

## Decision

TODO: record that we're going to use JSON and why

- At a high level, this can be separated into two different tasks: 
	1. generating test files, and 
	2. testing the test files

### Generating test files

TODO: explain what needs to be generated ... we need to make a chain of blocks
many functions already exist to do this as they are used in the Go unit tests.
We want to translate that into JSON blockchains ..
Note something about the difference between these functions and the real functions for making blocks ...
I think it mostly boils down to private keys ...

	- Generators use low-level test functions, manually configured in a way that outputs test cases, each representing a unique possible situation within Tendermint.
	- In case of lite client, a test case could look like this:

A light client test case begins with an initial state (a trusted header and validator set)
and takes a series of inputs in the form of what we call LiteBlock, which contains a header, commit, and validator set:
```
type TestCase struct {
    TestName       string      `json:"test_name"`
    Description    string      `json:"description"`
    Initial        Initial     `json:"initial"`	
    Input          []LiteBlock `json:"input"`
    ExpectedOutput string      `json:"expected_output"`	
}

type LiteBlock struct {
    SignedHeader     types.SignedHeader `json:"signed_header"`
    ValidatorSet     types.ValidatorSet `json:"validator_set"`
    NextValidatorSet types.ValidatorSet `json:"next_validator_set"`
}

type Initial struct {
    SignedHeader     types.SignedHeader `json:"signed_header"`
    NextValidatorSet types.ValidatorSet `json:"next_validator_set"`
    TrustingPeriod   time.Duration      `json:"trusting_period"`
    Now              time.Time          `json:"now"`
}
```

- TestName is a one word identifier for the test case that tells what functionality of the code we are testing.
- Description is a brief explanation of what the test case is i.e. what situation the case represents.
- Initial is the starting state or to say, a context for the test case. In terms of Lite Client, it can be the trusted state. In case of testing bisection, initial will have primary and alternative providers, trusted store, trust options, etc. which make the Client. ... TODO: explain this better / simpler
- Input is the actual data that is to be tested. In the case where we have multiple providers to fetch data from, input can be an array of array of LiteBlock i.e [][]LiteBlock, where [0][0]LiteBlock refers to the first LiteBlock of first provider.
- Expected output, for now, is a string that tells whether to expect an error or not. This will change once we have standardized error handling and should be error types or represent certain state.

A generator will look something like this: 
```
func CaseVerifyValidatorSetChangesLessThanOneThird(testCases *TestCases, valList ValList) {

	copyValList := valList.Copy() // To dereference pointers
	description := "Case: two lite blocks to fetch, less than 1/3 validator set changes, expects no error"

	initialNumOfVals := 4
	numOfValsToAdd := 1
	numOfValsToDelete := 1

	initial, input, _, _ := GenerateNextValsUpdateCase(copyValList, initialNumOfVals, numOfValsToAdd, numOfValsToDelete)
	testCase := GenerateTestCase(testName, description, initial, input, expectedOutputNoError)

	testCases.TC = append(testCases.TC, testCase)
}
```

The above function will output a test case where less than 1/3 of the validator set changes and it is then marhsalled to JSON.

* Writing tests: 

- Tests would work in this flow: 
    - Feed in relevant test cases from JSON files,
    - Set the context by copying over the initial state,
    - Build the test situation by fetching data from input and passing it on to native functions, that are being tested, to be able to fully simulate the condition and test its output
    - Check whether the expectation is satisfied by the output - passes if it does, otherwise fails. Output can be a return value or a state transition.

Since the levels of testing will vary, reason being we want maximum code coverage as well as maximum possible situations tested, the design of test functions will be different at each level. For example, in a unit test, it is more likely we will make assertions on the output of the function being tested, whereas in integration tests, we will be more interested in looking at state transitions.


## Status

Accepted

## Consequences


### Positive

Will help guide the development of Rust implementation. Also, be helpful to uncover possible bugs in the system and re-think implementation, if required.


