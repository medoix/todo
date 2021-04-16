package main

import (
	"encoding/json"
	"io"
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
}

func (c *Content) DelItem(lane, idx int) {
	c.Items[lane] = append(c.Items[lane][:idx], c.Items[lane][idx+1:]...)
}

func (c *Content) AddItem(lane, idx int, title string) {
	c.Items[lane] = append(c.Items[lane][:idx], append([]Item{Item{title, ""}}, c.Items[lane][idx:]...)...)
}

func (c *Content) Save(w io.Writer) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(c)
}
