package main

import (
	"testing"
)

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
