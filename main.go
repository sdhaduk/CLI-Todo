package main

import (
	"github.com/alecthomas/kong"
	"github.com/sdhaduk/CLI-Todo-List/cli"
)


var CLI struct {
	Add      cli.AddCmd      `cmd:"" help:"Add task to the list."`
	List     cli.ListCmd     `cmd:"" help:"Print all tasks in the list."`
	Delete   cli.DeleteCmd   `cmd:"" help:"Delete a task from the list."`
	Clear    cli.ClearCmd    `cmd:"" help:"Clear all tasks from the list."`
	Complete cli.CompleteCmd `cmd:"" help:"Mark a task as complete in the list."`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
