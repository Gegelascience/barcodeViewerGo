package main

import (
	"testing"
)

func TestIsCorrectEan(t *testing.T) {
	checkValue := "3666154117284"
	test := isCorrectEan(checkValue)
	if !test {
		t.Fatalf(`isCorrectEan("%v") should return true but have %t`, checkValue, test)
	}
	checkValue2 := "12345670"
	test2 := isCorrectEan(checkValue)
	if !test {
		t.Fatalf(`isCorrectEan("%v") should return true but have %t`, checkValue2, test2)
	}
}

func TestCalculateDigitCheck(t *testing.T) {
	checkValue := "366615411728"
	test := calculateDigitCheck(checkValue)
	if test != "4" {
		t.Fatalf(`calculateDigitCheck("%v") should return 4 but have %v`, checkValue, test)
	}
}
