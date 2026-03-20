package cli

import (
	"errors"
	"fmt"
	"github.com/sdhaduk/CLI-Todo-List/internal/todo"
)

type AddCmd struct {
	Task string `arg:"" name:"task" help:"Task to be added to the list."`
}

type ListCmd struct {}

type DeleteCmd struct {
	TaskId int `arg:"" name:"number" help:"Task ID in list."`
}

type ClearCmd struct {}

type CompleteCmd struct {
	TaskId int `arg:"" name:"number" help:"Task ID in list."`
}

func (a *AddCmd) Run() error {
	err := todo.Add(a.Task)
	if err != nil {
		if errors.Is(err, todo.InvalidName) {
			fmt.Println("Task name is not valid.")
			return nil
		}
		fmt.Println("Problem occured while trying to add the task...")
		return err
	}
	fmt.Println("Sucessfully added task.")
	return nil
}

func (l *ListCmd) Run() error {
	msg, err := todo.List()
	if err != nil {
		fmt.Println("Problem occurred while trying to list items...")
		return err
	}
	fmt.Println(msg)
	return nil
}

func (d *DeleteCmd) Run() error {
	err := todo.Delete(d.TaskId)
	if err != nil {
		fmt.Println("Problem occurred while trying to delete...")
		return err
	}
	fmt.Println("Sucessfully deleted task ", d.TaskId, ".")
	return nil
}

func (c *ClearCmd) Run() error {
	err := todo.Clear()
	if err != nil {
		fmt.Println("Problem occurred while trying to clear list.")
		return err
	}

	fmt.Println("Successfully cleared list.")
	return nil
}

func (c *CompleteCmd) Run() error {
	err := todo.Complete(c.TaskId)
	if err != nil {
		fmt.Println("Problem occured while trying to mark task ", c.TaskId, " as completed...")
		return err
	}
	fmt.Println("Sucessfully marked task ", c.TaskId, " as completed.")
	return nil
}