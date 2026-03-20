package todo

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var InvalidName = errors.New("Invalid task name")

func List() (string, error) {
	tasks, err := readFromFile(filePath)
	if err != nil {
		return "", err
	}
	if len(tasks) == 0 {
		msg := "Your todo list is empty!" 
		return msg, nil 
	}

	var output strings.Builder
	for _, t := range tasks {
		status := " "
		if t.Completed {
			status = "✓"
		}
		line := fmt.Sprintf("[%s] %d: %s\n", status, t.Id, t.Name)
		output.WriteString(line)
	}
	return output.String(), nil
}

func Add(name string) error {
	var id int

	if len(name) == 0 {
		return InvalidName
	}

	if !utf8.ValidString(name) {
		return InvalidName
	}
	
	tasks, err := readFromFile(filePath)
	if err != nil {
		return err
	}
	
	if len(tasks) == 0 {
		id = 1
	} else {
		for _, task := range tasks {
			id = max(id, task.Id)
		}
	}

	newTask := Task{
		Id: id + 1,
		Name: name,
		Completed: false,
	}
	tasks = append(tasks, newTask)

	err = saveToFile(filePath, tasks)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	tasks, err := readFromFile(filePath)
	if err != nil {
		return err
	}
	newTasks := []Task{}

	for _, task := range tasks {
		if task.Id != id {
			newTasks = append(newTasks, task)
		}
	}

	err = saveToFile(filePath, newTasks)
	if err != nil {
		return err
	}
	return nil
}

func Complete(id int) error {
	tasks, err := readFromFile(filePath)
	if err != nil {
		return err
	}

	for index, task := range tasks {
		if task.Id == id {
			tasks[index].Completed = true
		}
	}

	err = saveToFile(filePath, tasks)
	if err != nil {
		return err
	}
	return nil
}

func Clear() error {
	return clearFile(filePath)
}