package cli

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

