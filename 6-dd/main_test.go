package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// test funcs
func createfile(data []byte, filename string) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func openfile(filename string) (content []byte) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return
}

type TestCaseType struct {
	testData   []byte
	testOffset int64
	testLimit  int64
	testResult []byte
}

func TestCopyfile(t *testing.T) {

	in := "in"
	out := "out"

	var testCases = []TestCaseType{
		{testData: []byte("1234567890"), testOffset: 0, testLimit: 0, testResult: []byte("1234567890")},
		{testData: []byte("1234567890"), testOffset: 0, testLimit: 5, testResult: []byte("12345")},
		{testData: []byte("1234567890"), testOffset: 2, testLimit: 5, testResult: []byte("34567")},
		{testData: []byte("1234567890"), testOffset: 8, testLimit: 5, testResult: []byte("90")},
		{testData: []byte("1234567890"), testOffset: 8, testLimit: 0, testResult: []byte("90")},
	}

	for _, testCase := range testCases {
		createfile(testCase.testData, in)
		copyfile(in, out, testCase.testOffset, testCase.testLimit)
		require.Equal(t, openfile(out), testCase.testResult)
	}
	os.Remove(in)
	os.Remove(out)
}
