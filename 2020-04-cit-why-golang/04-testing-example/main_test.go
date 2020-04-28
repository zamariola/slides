package main

import "testing"

func TestShouldSumPositiveValues(t *testing.T) {
	expectedValue := 3
	actual := sumPositiveValues(1, 2)

	if actual != expectedValue {
		t.Errorf("Wrong sum value, expected %d, got: %d", expectedValue, actual)
	}
}

func TestShouldNotSumNegativeValues(t *testing.T) {
	expectedValue := 0
	actual := sumPositiveValues(-1, -2)

	if actual != expectedValue {
		t.Errorf("Wrong sum value, expected %d, got: %d", expectedValue, actual)
	}
}

func TestShouldNotSumOneNegativeValue(t *testing.T) {
	expectedValue := 0
	actual := sumPositiveValues(1, -2)

	if actual != expectedValue {
		t.Errorf("Wrong sum value, expected %d, got: %d", expectedValue, actual)
	}
}
