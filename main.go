package main

import (
	"os"

	"github.com/o98k-ok/command/internel/task"
	"github.com/o98k-ok/lazy/v2/alfred"
)

func main() {
	app := alfred.NewApp("one-punch client")
	app.Bind("edit", task.EditTask)
	app.Bind("search", task.SearchTask)
	app.Run(os.Args)
}
