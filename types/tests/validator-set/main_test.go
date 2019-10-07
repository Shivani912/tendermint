package main

import (
	"testing"
	"os"
	"fmt"
	"regexp"
)


func TestCase(t *testing.T) {

	mode := os.Args[1]
	
	switch mode {
	case "positive":
		t.Run("Positive Case", PostiveCase)
	case "negative":
		t.Run("Negative Case", NegativeCase)
	}
	
}

func PostiveCase(t *testing.T) {

	for i := 2; ;i++ {
		
		folder := getFolderFromCommandLineArgs(i)
	
		if folder == "none" {
			break
		}

		input, expectedOutput, err := getData(folder)
		if err == 1 {
			fmt.Println("Can't find folder: ", folder)
			break
		}

		if output := string(toHex(input.Hash())); !(output == expectedOutput) {
			t.Errorf("Test failed: %v inputted, %s expected, received: %s", input, expectedOutput, output)

		}
	}
	

}


func NegativeCase(t *testing.T) {
	for i := 2; ;i++ {
		
		folder := getFolderFromCommandLineArgs(i)
	
		if folder == "none" {
			break
		}

		input, expectedOutput, err := getData(folder)
		if err == 1 {
			fmt.Println("Can't find folder: ", folder)
			break
		}
		if output := string(toHex(input.Hash())); output == expectedOutput {
			t.Errorf("Test failed: %v inputted, %s expected, received: %s", input, expectedOutput, output)

		}
	}

}

func getFolderFromCommandLineArgs(i int) string {
	folder := os.Args[i]

	r, _ := regexp.Compile("-test.timeout=")
	match:= r.FindString(folder)
	// fmt.Println(match)
	if match != "" {
		return "none"
	}
	return folder
}
