package main

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// test funcs
func jobWithError() error {
	fmt.Println("hello error")
	return errors.New("error1")
}

func jobWithSleep() error {
	time.Sleep(1 * time.Second)
	fmt.Println("sleep")
	return nil
}

type TestCaseType struct {
	testFuncs      []func() error
	testThreads    int
	testMaxErrors  int
	expectedResult error
}

func TestRun(t *testing.T) {

	//create test slice of functions
	var fns []func() error
	for i := 0; i < 5; i++ {
		fns = append(fns, jobWithError)
	}
	for i := 0; i < 5; i++ {
		fns = append(fns, jobWithSleep)
	}

	var testCases = []TestCaseType{
		{testFuncs: fns, testThreads: 5, testMaxErrors: 6, expectedResult: nil},
		{testFuncs: fns, testThreads: 5, testMaxErrors: 5, expectedResult: errors.New("Max errors")},
	}

	for _, testCase := range testCases {
		require.Equal(t, Run(testCase.testFuncs, testCase.testThreads, testCase.testMaxErrors), testCase.expectedResult)
	}

}
