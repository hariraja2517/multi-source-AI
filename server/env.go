package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func loadEnv(fileName string) (map[string]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`^(\w+)=(.*)$`)
	envMap := make(map[string]string)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := pattern.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		key, value := match[1], match[2]
		envMap[key] = value
	}

	return envMap, nil
}
