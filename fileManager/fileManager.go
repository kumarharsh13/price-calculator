package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		file.Close()
		return nil, errors.New("Error while opening the file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Error while parsing the file")
	}

	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		file.Close()
		return errors.New("Failed to create the file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	file.Close()
	return nil
}
