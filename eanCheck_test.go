package main

import (
	"testing"
)

func TestICorrectEan_noerror(t *testing.T) {
	checkValue := "3666154117284"
	test := isCorrectEan(checkValue)
	if !test {
		t.Fatalf(`isCorrectEan("%v") should return true but have %t`, checkValue, test)
	}
}
