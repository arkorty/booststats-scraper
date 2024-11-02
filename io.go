package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
)

func appendToJSONFile(profile Profile, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(profile); err != nil {
		return err
	}

	return nil
}

func readAssignments(filename string) ([]string, error) {
	var assignments []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			assignments = append(assignments, trimmed)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}
