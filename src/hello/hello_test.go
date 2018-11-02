package main

import "testing"

func TestThreeNeighbours (t *testing.T) {

	var input [2][2] string;
	input[0] = [2]string{"*", "*"}
	input[1] = [2]string{"*", "."}

	output := Sum(2, 2,  input)
	expectedOutput := "****"

	if(output != expectedOutput) {
		t.Errorf("Output was incorrect, got: %s, wanted: %s", output, expectedOutput);
	}
}