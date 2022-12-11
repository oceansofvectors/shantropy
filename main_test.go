package main

import (
	"io/ioutil"
	"math"
	"os"
	"testing"
)

func TestFindImportantStrings(t *testing.T) {
	// Test some simple strings with known entropies
	testCases := []struct {
		s          string
		minEntropy float64
		expected   bool
	}{
		{"", 0.0, false},
		{"a", 0.0, true},
		{"ab", 1.0, true},
		{"abcdefghijklmnopqrstuvwxyz", 4.700439718141092, false},
		{"hello, world!", 3.180327868852459, true},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 0.0, true},
	}
	for _, tc := range testCases {
		found := FindImportantStrings(tc.s, "", tc.minEntropy)
		if found != tc.expected {
			t.Errorf("expected FindImportantStrings of %q to be %v, got %v", tc.s, tc.expected, found)
		}
	}
}

func TestShannonEntropy(t *testing.T) {
	// Test some simple strings with known entropies
	testCases := []struct {
		s         string
		entropy   float64
		tolerance float64
	}{
		{"", 0.0, 0.00001},
		{"a", 0.0, 0.00001},
		{"ab", 1.0, 0.00001},
		{"abcdefghijklmnopqrstuvwxyz", 4.700439718141092, 0.00001},
		{"hello, world!", 3.180832987205441, 0.00001},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 0.0, 0.00001},
	}
	for _, tc := range testCases {
		entropy := ShannonEntropy(tc.s)
		if math.Abs(entropy-tc.entropy) > tc.tolerance {
			t.Errorf("expected entropy of %q to be %v, got %v", tc.s, tc.entropy, entropy)
		}
	}
}

func TestReadFile(t *testing.T) {
	// Create a temporary file for testing
	file, err := ioutil.TempFile("", "test-file")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(file.Name())

	// Write some test data to the file
	data := "Hello, World!"
	if _, err := file.Write([]byte(data)); err != nil {
		t.Error(err)
	}

	// Use the ReadFile function to read the data back
	contents, err := ReadFile(file.Name())
	if err != nil {
		t.Error(err)
	}

	// Check that the contents of the file are correct
	if contents != data {
		t.Errorf("expected contents to be %q, got %q", data, contents)
	}
}
