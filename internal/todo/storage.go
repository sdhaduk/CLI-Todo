package todo 

import (
	"errors"
	"os"
	"encoding/json"
	"io"
)

const filePath = "todo_list.json"

func ensureFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return createFile(filePath)
	}	
	return err
}

func createFile(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func readFromFile(filePath string) ([]Task, error) {
	err := ensureFileExists(filePath)
	if err != nil {
		return nil, err
	}
	
	f, err := os.Open(filePath)
	if err != nil {		
		return nil, err
	}
	defer f.Close()
	
	var tasks []Task
	err = json.NewDecoder(f).Decode(&tasks)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return []Task{}, nil
		}
		return nil, err
	}
	return tasks, nil
}

func saveToFile(filePath string, tasks []Task) error {
	err := ensureFileExists(filePath)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}