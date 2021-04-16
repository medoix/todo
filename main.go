package main

import (
	"log"
	"os"

	"github.com/rivo/tview"
)

func get_path() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	home, _ := os.UserHomeDir()
	os.Mkdir(home+"/.todo", os.ModePerm)

	return home + "/.todo/todo"
}

func main() {
	path := get_path()
	var content *Content
	f, err := os.Open(path)
	if err == nil {
		content = NewContentIo(f)
		f.Close()
	}

	if content == nil {
		content = NewContentDefault()
	}

	app := tview.NewApplication()

	lanes := NewLanes(content, app)

	app.SetRoot(lanes.GetUi(), true)

	if err := app.Run(); err != nil {
		log.Fatal("Error running application: %s\n", err)
	}

	f, err = os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	content.Save(f)
	f.Close()
}
