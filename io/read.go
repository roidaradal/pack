package io

import (
	"encoding/json"
	"os"
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
