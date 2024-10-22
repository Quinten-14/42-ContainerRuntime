package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

// Function to load the config file
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

// Function to open the config file in Vim
func editConfigFile(filepath string) error {
	cmd := exec.Command("vim", filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func uiDrawer() {
	// ASCII art banner with colors
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
	fmt.Printf("%s1. Start All Tests%s\n", Green, Reset)
	fmt.Printf("%s2. Start Specific Test%s\n", Yellow, Reset)
	fmt.Printf("%s3. Edit Config File%s\n", Blue, Reset)
	fmt.Printf("%s4. Exit%s\n\n", Red, Reset)
	fmt.Print(Bold + "Choose an option: " + Reset)

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println(Green + "Starting all tests...\n" + Reset)
		// startAllTests(config)
	case 2:
		fmt.Println(Yellow + "Starting specific test...\n" + Reset)
		// startSpecificTest(config)
	case 3:
		fmt.Println(Blue + "Editing config file...\n" + Reset)
		err := editConfigFile("config.json")
		if err != nil {
			fmt.Println(Red + "Failed to edit config file: " + err.Error() + Reset)
			return
		}
		// Reload the config after editing
		config, err = loadConfig("config.json")
		if err != nil {
			fmt.Println(Red + "Failed to reload config: " + err.Error() + Reset)
		} else {
			fmt.Println(Green + "Config reloaded successfully." + Reset)
		}
	case 4:
		fmt.Println(Red + "Exiting...\n" + Reset)
		os.Exit(0)
	default:
		fmt.Println(Red + "Invalid choice, please try again.\n" + Reset)
	}
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Printf(Red+"Error loading config: %v\n"+Reset, err)
		return
	}

	uiDrawer()

	for {
		input(config)
	}
}

