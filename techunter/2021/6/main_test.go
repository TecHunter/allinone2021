package main

import (
	"testing"
)

const (
	FirstCaseResult  = 5934
	SecondCaseResult = 26984457539
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
