package main

import (
	"testing"
)

const (
	FirstCaseResult  = 37
	SecondCaseResult = 168
)

func TestFirstCase(t *testing.T) {

	ret := FirstCase(".test")
	if ret != FirstCaseResult {
		t.Errorf("First case was incorrect, got %d want %d", ret, FirstCaseResult)
	}

}

func TestSecondCase(t *testing.T) {

	ret := SecondCase(".test")
	if ret != SecondCaseResult {
		t.Errorf("Second case was incorrect, got %d want %d", ret, SecondCaseResult)
	}

}

func TestSum(t *testing.T) {
	ret := Sum(5)
	if ret != 15 {
		t.Errorf("Expected result of Sum(5) is %d, found %d", 15, ret)
	}
	ret = Sum(16 - 5)
	if ret != 66 {
		t.Errorf("Expected result of Sum(5) is %d, found %d", 66, ret)
	}
}

func TestMean(t *testing.T) {
	input := ProcessInput(".test")
	_, mean := Mean(input)

	if mean != 5 {
		t.Errorf("Mean function was incorrect, got %d want %d", mean, 5)
	}
}
