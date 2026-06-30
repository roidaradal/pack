package io

import (
	"encoding/json"
	"os"
	"strings"
)

// ReadJSON reads a JSON object of given type from given path
func ReadJSON[T any](path string) (T, error) {
	var item T
	bytes, err := os.ReadFile(path)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// ReadJSONList reads a JSON list of given type from given path
func ReadJSONList[T any](path string) ([]T, error) {
	return ReadJSON[[]T](path)
}

// ReadJSONMap reads a JSON map of given value type from given path
func ReadJSONMap[V any](path string) (map[string]V, error) {
	return ReadJSON[map[string]V](path)
}

// ReadFile reads the string contents of given path
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadRawLines reads the lines of given path
func ReadRawLines(path string) ([]string, error) {
	text, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(text, "\n"), nil
}

// ReadLines reads the lines of given path, and each line is trimmed for whitespace
func ReadLines(path string) ([]string, error) {
	lines, err := ReadRawLines(path)
	if err != nil {
		return nil, err
	}
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines, nil
}
