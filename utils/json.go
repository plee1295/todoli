package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ReadFromJSON reads data from a JSON file and unmarshals it into the provided target.
func ReadFromJSON(filename string, target interface{}) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, treat as empty.
		}
		return fmt.Errorf("error reading file: %w", err)
	}

	if err := json.Unmarshal(file, target); err != nil {
		return fmt.Errorf("error unmarshalling file: %w", err)
	}

	return nil
}

// WriteToJSON marshals the provided data and writes it to a JSON file.
func WriteToJSON(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

// ReadInput reads input from the user with an optional default value.
func ReadInput(prompt string, defaultValue string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	if defaultValue != "" {
		fmt.Printf("%s (%s): ", prompt, defaultValue)
	} else {
		fmt.Printf("%s: ", prompt)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue, nil
	}
	return input, nil
}
