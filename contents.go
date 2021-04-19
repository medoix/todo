package main

import (
	"encoding/json"
	"io"
	"os"
	"log"
)

type Item struct {
	Title string
	Note  string
}

type Content struct {
	Titles []string
	Items  [][]Item
}

func NewContentIo(r io.Reader) *Content {
	decoder := json.NewDecoder(r)
	c := &Content{}
	if err := decoder.Decode(c); err != nil {
		return nil
	}
	return c
}

func NewContentDefault() *Content {
	ret := &Content{}
	ret.Titles = []string{"To Do", "In Progress", "Done"}
	ret.Items = make([][]Item, 3)
	return ret
}

func (c *Content) GetNumLanes() int {
	return len(c.Titles)
}

func (c *Content) GetLaneTitle(idx int) string {
	return c.Titles[idx]
}

func (c *Content) GetLaneItems(idx int) []Item {
	return c.Items[idx]
}

func (c *Content) MoveItem(fromlane, fromidx, tolane, toidx int) {
	item := c.Items[fromlane][fromidx]
	// https://github.com/golang/go/wiki/SliceTricks
	c.Items[fromlane] = append(c.Items[fromlane][:fromidx], c.Items[fromlane][fromidx+1:]...)
	c.Items[tolane] = append(c.Items[tolane][:toidx], append([]Item{item}, c.Items[tolane][toidx:]...)...)
	saveToDo(todoPath())
}

func todoPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	home, _ := os.UserHomeDir()
	os.Mkdir(home+"/.todo", os.ModePerm)
	return home + "/.todo/todo"
}

/*
TODO: Look into moving this variable into below function
and use pointers / dereference
*/
var content *Content

func openToDo(path string) {
	f, err := os.Open(path)
	if err == nil {
		content = NewContentIo(f)
		f.Close()
	}

	if content == nil {
		content = NewContentDefault()
	}
}

func saveToDo(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	content.Save(f)
	f.Close()
}

func (c *Content) DelItem(lane, idx int) {
	c.Items[lane] = append(c.Items[lane][:idx], c.Items[lane][idx+1:]...)
	saveToDo(todoPath())
}

func (c *Content) AddItem(lane, idx int, title string) {
	c.Items[lane] = append(c.Items[lane][:idx], append([]Item{Item{title, ""}}, c.Items[lane][idx:]...)...)
	saveToDo(todoPath())
}

func (c *Content) Save(w io.Writer) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(c)
}
