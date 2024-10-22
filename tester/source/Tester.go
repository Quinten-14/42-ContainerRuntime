package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"bufio"
)

type Test struct {
    Index   int    `json:"index"`
	Name    string `json:"name"`
	Command string `json:"command"`
    Expected string `json:"expected"`
}

type Config struct {
	MinishellPath   string   `json:"minishell_path"`
	TestsFile       string   `json:"tests_file"`
	LogsPath        string   `json:"logs_path"`
	TestTimeout     int      `json:"test_timeout"`
	CustomTests     bool     `json:"custom_tests"`
	AdditionalTests []string `json:"additional_tests"`
}

// Color constants for the terminal
const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Purple     = "\033[35m"
	Cyan       = "\033[36m"
	White      = "\033[37m"
	Bold       = "\033[1m"
	BgBlue     = "\033[44m"
)

func loadTests(filename string) ([]Test, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var tests []Test
	err = json.Unmarshal(file, &tests)
	if err != nil {
		return nil, err
	}
	return tests, nil
}

func replaceEnvVariables(input string) string {
	words := strings.Fields(input)

	for i, word := range words {
		if strings.HasPrefix(word, "$(") && strings.HasSuffix(word, ")") {
			varName := word[2 : len(word)-1]

			envValue := os.Getenv(varName)
			if envValue != "" {
				words[i] = envValue 
			}
		}
	}

	return strings.Join(words, " ")
}

func runTest(minishellPath string, test Test) error {
	cmd := exec.Command(minishellPath)

	cmd.Stdin = strings.NewReader(test.Command + "\nexit\n")

	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}

	outputStr := string(outputBytes)

	outputLines := strings.Split(outputStr, "\n")

	var filteredOutputLines []string
	for _, line := range outputLines {
		trimmedLine := strings.TrimSpace(line)

		if strings.Contains(trimmedLine, "Minishell-42:") {
			continue
		}

		if trimmedLine == test.Command {
			continue
		}

		if strings.Contains(trimmedLine, "exit") {
			continue
		}

		if trimmedLine == "" {
			continue
		}

		filteredOutputLines = append(filteredOutputLines, trimmedLine)
	}

	finalOutput := strings.Join(filteredOutputLines, "\n")
	finalOutput = strings.TrimSpace(finalOutput)

	expectedStr := replaceEnvVariables(test.Expected)
	expectedStr = strings.TrimSpace(expectedStr)

	if finalOutput == expectedStr {
		fmt.Printf("%sTest %d (%s): Passed%s\n", Green, test.Index, test.Name, Reset)
	} else {
		fmt.Printf("%sTest %d (%s): Failed\nExpected: %s\nGot: %s%s\n", Red, test.Index, test.Name, expectedStr, finalOutput, Reset)
	}

	return nil
}

func runAllTests(config *Config) error {
	tests, err := loadTests(config.TestsFile)
	if err != nil {
		return fmt.Errorf("failed to load tests: %v", err)
	}

	for _, test := range tests {
		err := runTest(config.MinishellPath, test)
		if err != nil {
			fmt.Println(Red + err.Error() + Reset)
		}
	}

	return nil
}

func runSelectedTests(config *Config, selectedIndices []int) error {
	tests, err := loadTests(config.TestsFile)
	if err != nil {
		return fmt.Errorf("failed to load tests: %v", err)
	}

	for _, index := range selectedIndices {
		if index > 0 && index <= len(tests) {
			err := runTest(config.MinishellPath, tests[index-1]) // Test indices are 1-based
			if err != nil {
				fmt.Println(Red + err.Error() + Reset)
			}
		} else {
			fmt.Printf("%sInvalid test index: %d%s\n", Red, index, Reset)
		}
	}

	return nil
}


func loadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func editConfigFile(filepath string) error {
	cmd := exec.Command("vim", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func uiDrawer() {
	banner := fmt.Sprintf(`
%s%s
___  ___              _          _ _   _____         _            
|  \/  |             | |        | | | |_   _|       | |           
| .  . |_ _ __  _ ___| |__   ___| | |   | | ___  ___| |_ ___ _ __ 
| |\/| | | '_ \| / __| '_ \ / _ \ | |   | |/ _ \/ __| __/ _ \ '__|
| |  | | | | | | \__ \ | | |  __/ | |   | |  __/\__ \ ||  __/ |   
\_|  |_/_|_| |_|_|___/_| |_|\___|_|_|   \_/\___||___/\__\___|_|   
%s===============================================================%s
                Part of the Minishell Toolkit
===============================================================%s
`, Cyan, Bold, Reset, Cyan, Reset)

	fmt.Println(banner)
}

func input(config *Config) {
	fmt.Printf("\n%s1. Start All Tests%s\n", Green, Reset)
	fmt.Printf("%s2. Start Specific Test%s\n", Yellow, Reset)
	fmt.Printf("%s3. Edit Config File%s\n", Blue, Reset)
	fmt.Printf("%s4. Exit%s\n\n", Red, Reset)
	fmt.Print(Bold + "Choose an option (1-4): " + Reset)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	if len(choice) != 1 || choice < "1" || choice > "4" {
		fmt.Println(Red + "Invalid choice, please enter a single digit between 1 and 4." + Reset)
		return
	}

	switch choice {
	case "1":
		fmt.Println(Green + "Starting all tests...\n" + Reset)
		err := runAllTests(config)
		if err != nil {
			fmt.Println(Red + err.Error() + Reset)
		}
		fmt.Println(Green + "Finished Testing\n" + Reset)
	case "2":
		fmt.Println(Yellow + "Starting specific test...\n" + Reset)
		fmt.Println(Yellow + "Input Tests like this -> 1,2 (Single digit only, No Spaces Allowed)\n", Reset)
		var indicesInput string
		fmt.Scanln(&indicesInput)

		indicesStr := strings.Split(indicesInput, ",")
		var indices []int
		for _, str := range indicesStr {
			index, err := strconv.Atoi(strings.TrimSpace(str))
			if err == nil {
				indices = append(indices, index)
			}
		}

		err := runSelectedTests(config, indices)
		if err != nil {
			fmt.Println(Red + err.Error() + Reset)
		}
	case "3":
		fmt.Println(Blue + "Editing config file...\n" + Reset)
		err := editConfigFile("source/config.json")
		if err != nil {
			fmt.Println(Red + "Failed to edit config file: " + err.Error() + Reset)
			return
		}
		config, err = loadConfig("source/config.json")
		if err != nil {
			fmt.Println(Red + "Failed to reload config: " + err.Error() + Reset)
		} else {
			fmt.Println(Green + "Config reloaded successfully." + Reset)
		}
	case "4":
		fmt.Println(Red + "Exiting...\n" + Reset)
		os.Exit(0)
	}
}

func main() {
	config, err := loadConfig("source/config.json")
	if err != nil {
		fmt.Printf(Red+"Error loading config: %v\n"+Reset, err)
		return
	}

	uiDrawer()

	for {
		input(config)
	}
}

