package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TestResult holds the name and status of a test.
type TestResult struct {
	Name   string
	Status string
}

// Colors
const (
	red = "\033[31m"
	green = "\033[32m"
	reset = "\033[0m"
)

// Reads the Logs
func ReadTestResults(filename string) (map[string]TestResult, error) {
	results := make(map[string]TestResult)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ".")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		testName := strings.TrimSpace(parts[0])
		testStatus := strings.TrimSpace(parts[1])
		results[testName] = TestResult{Name: testName, Status: testStatus}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func ColorStatus(status string) string {
		if status == "ok" {
			return green + status + reset
		}
		return red + status + reset
}

// Compares the old results and the new results
func CompareTestResults(oldResults, newResults map[string]TestResult) {
	for testName, oldResult := range oldResults {
		newResult, exists := newResults[testName]

		if !exists {
			fmt.Printf("%s: %s (missing in new results)\n", oldResult.Name, oldResult.Status)
			continue
		}

		fmt.Printf("%s: [%s] [%s]", oldResult.Name, ColorStatus(oldResult.Status), ColorStatus(newResult.Status))

		if newResult.Status == "ok" && oldResult.Status == "ko" {
			fmt.Println(green + " -> Improved" + reset)
		} else if newResult.Status == oldResult.Status {
			fmt.Println(" -> No change")
		} else {
			fmt.Println(red + " -> Degraded" + reset)
		}
	}

	for testName, newResult := range newResults {
		if _, exists := oldResults[testName]; !exists {
			fmt.Printf("%s: %s (new results)\n", newResult.Name, newResult.Status)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run comparer.go <old_logs> <new_logs>")
		return
	}

	oldResultsFile := os.Args[1]
	newResultsFile := os.Args[2]

	oldResults, err := ReadTestResults(oldResultsFile)
	if err != nil {
		fmt.Printf("Error reading old results: %v\n", err)
		return
	}

	newResults, err := ReadTestResults(newResultsFile)
	if err != nil {
		fmt.Printf("Error reading new results: %v\n", err)
		return
	}

	CompareTestResults(oldResults, newResults)
}
