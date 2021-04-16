# To Do

A simple Kanban style To Do / Task list for your terminal.

Stores data in a simple JSON document in `$HOME/.todo/todo`.

## Keys / Shortcuts

* a: Add item
* d: Delete item
* e: Edit item
* n: View/edit notes for item
* q: Quit
* Enter: Select / deselect item (selected items can be moved)
* j/k or Up/Down Arrows: Move item up or down
* h/l or Left/Right Arrows: Move item between columns
* Tab: Move in forms

## Improvements

- [X] Create config system

  - [ ] Allow rows OR columns

  - [ ] Define Note editor program

## Install

### Golang

```bash
go get gitlab.com/medoix/todo
go run todo
```

### Options

```bash
todo [path]

path: Location to store or retrieve a todo list, if not provided `~/.todo/todo` will be used.
 ```
