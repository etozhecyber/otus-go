package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCaseType struct {
	testData       string
	expectedResult string
}

func TestUnpack(t *testing.T) {

	var testCases = []TestCaseType{
		{testData: "a4bc2d5e2", expectedResult: "aaaabccdddddee"},
		{testData: "abcd", expectedResult: "abcd"},
		{testData: "45", expectedResult: ""},
		{testData: "a123", expectedResult: "a"},
		{testData: "a0123", expectedResult: ""},
		{testData: "a0a", expectedResult: "a"},
		{testData: "0a", expectedResult: "a"},
		{testData: "0", expectedResult: ""},
		{testData: "a", expectedResult: "a"},
		{testData: "01a2c1", expectedResult: "aac"},
	}

	for _, testCase := range testCases {
		require.Equal(t, unpack(testCase.testData), testCase.expectedResult)
	}

}
