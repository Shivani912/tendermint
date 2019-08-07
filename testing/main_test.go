package main

import (
	"testing"
)

// TestCase1 and TestCase2 are positive test cases

func TestCase1(t *testing.T) {
	folder := "./testCase1/"
	input, expectedOutput := getData(folder)

	if output := string(toHex(input.Hash())); !(output == expectedOutput) {
		t.Errorf("Test failed: %v inputted, %s expected, received: %s", input, expectedOutput, output)

	}

}

func TestCase2(t *testing.T) {
	folder := "./testCase2/"
	input, expectedOutput := getData(folder)

	if output := string(toHex(input.Hash())); !(output == expectedOutput) {
		t.Errorf("Test failed: %v inputted, %s expected, received: %s", input, expectedOutput, output)

	}

}

// TestCase3 is a negative test case
// the merkle root in file is incorrect
// therefore output should never be equal to expectedOutput
func TestCase3(t *testing.T) {
	folder := "./testCase3/"
	input, expectedOutput := getData(folder)

	if output := string(toHex(input.Hash())); output == expectedOutput {
		t.Errorf("Test failed: %v inputted, %s expected, received: %s", input, expectedOutput, output)

	}

}
