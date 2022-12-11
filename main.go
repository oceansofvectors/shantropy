package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// ShannonEntropy calculates the Shannon entropy of a string,
// which is a measure of the uncertainty of the string.
func ShannonEntropy(s string) float64 {
	// Create a map to count the occurrences of each
	// character in the string.
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}

	// Calculate the probability of each character in the
	// string, and sum the entropies of each character.
	entropy := 0.0
	for _, count := range counts {
		p := float64(count) / float64(len(s))
		entropy += p * math.Log2(p)
	}

	// Return the negative of the sum of the entropies,
	// as the Shannon entropy is defined as the negative
	// of the sum of the probabilities of each character
	// multiplied by the log of the probabilities.
	return -entropy
}

// FindImportantStrings searches the input string for substrings
// with a high Shannon entropy, which may indicate that they
// are important or interesting.
func FindImportantStrings(s string, filepath string, minEntropy float64) bool {
	// Split the input string into substrings of varying
	// lengths, and calculate the Shannon entropy of each
	// substring.
	found := false
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			entropy := ShannonEntropy(substring)
			if entropy >= minEntropy {
				found = true
			}
		}
	}
	return found
}

// ReadFile reads the contents of the file at the given path
// and returns the contents as a string.
func ReadFile(filePath string) (string, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file contents and return them as a string
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func main() {
	// Get the directory path and minimum entropy from the command line arguments
	dirPath := os.Args[1]
	minEntropy, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over all files in the given directory, including subdirectories
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Read the file and do something with its contents
		if !info.IsDir() {
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// Find important strings in the file contents
			found := FindImportantStrings(string(contents), path, minEntropy)
			if found {
				fmt.Println(path)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
