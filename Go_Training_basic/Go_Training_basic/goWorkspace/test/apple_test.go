package main

import "testing"

func TestMean(t *testing.T) {

	var Array []int = []int{5, 6, 4}
	testVar := mean(Array)
	if testVar != 5 {
		t.Errorf("its an Error !!!")
	}

}
