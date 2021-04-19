package main

import (
	"log"

	"github.com/rivo/tview"
)

func main() {
	openToDo(todoPath())
	app := tview.NewApplication()
	lanes := NewLanes(content, app)
	app.SetRoot(lanes.GetUi(), true)
	if err := app.Run(); err != nil {
		log.Fatal("Error running application: %s\n", err)
	}
	saveToDo(todoPath())
}
