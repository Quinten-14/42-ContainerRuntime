package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Test struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	Expected string `json:"expected"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> <path to minishell> <path to tests.json>")
		return
	}

	minishellPath := os.Args[1]
	testsFilePath := os.Args[2]

	tests, err := loadTests(testsFilePath)
	if err != nil {
		fmt.Printf("Error loading tests: %v\n", err)
		return
	}

	for _, test := range tests {
		result, err := runMinishellCommand(minishellPath, test.Command)
		if err != nil {
			fmt.Printf("%s -> ko (%s)\n", test.Name, err)
		} else if strings.TrimSpace(result) == strings.TrimSpace(test.Expected) {
			fmt.Printf("%s -> ok\n", test.Name)
		} else {
			fmt.Printf("%s -> ko (expected: %s, got: %s)\n", test.Name, test.Expected, result)
		}
	}
}

func loadTests(filename string) ([]Test, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tests []Test
	if err := json.Unmarshal(file, &tests); err != nil {
		return nil, err
	}

	return tests, nil
}

func runMinishellCommand(minishellPath, cmd string) (string, error) {
	cmdExec := exec.Command(minishellPath)
	stdin, err := cmdExec.StdinPipe()
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	cmdExec.Stdout = &out

	if err := cmdExec.Start(); err != nil {
		return "", err
	}

	if _, err := stdin.Write([]byte(cmd + "\n")); err != nil {
		return "", err
	}

	stdin.Close()
	cmdExec.Wait()

	return out.String(), nil
}

