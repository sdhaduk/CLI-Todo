package todo 

import (
	"strings"
	"fmt"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

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
	tasks, err := readFromFile(filePath)
	if err != nil {
		return err
	}
	
	var id int
	
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

