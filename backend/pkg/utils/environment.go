package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Environment() error {
	// Reading the .env file
	content, err := os.ReadFile(".env")
	if err != nil {
		return errors.New("cannot read.env file")
	}

	// Splitting the file into lines
	lines := bufio.NewScanner(strings.NewReader(string(content)))
	for lines.Scan() {
		line := lines.Text()

		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Interpreting each line
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			// Set the environment variable
			err := os.Setenv(key, value)
			if err != nil {
				return errors.New("Cannot set environment variable: " + key + " = " + value)
			}
		} else {
			return errors.New("Invalid line in .env file: " + line)
		}
	}

	if err := lines.Err(); err != nil {
		return errors.New("cannot read .env file")
	}
	return nil
}
