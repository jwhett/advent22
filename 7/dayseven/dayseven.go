package dayseven

import (
	"fmt"
	"strings"

	h "github.com/jwhett/advent22/helpers"
)

type ParserAction int8

// The result of Parse() is one of these actions.
const (
	ParserCommandNotFound = -1
	ParserCdRoot          = iota
	ParserCdParent
	ParserCdChild
	ParserListCwd
	ParserFoundDir
	ParserFoundFile
)

func Parse(line string) ParserAction {
	if len(line) == 0 {
		// No commands on empty lines.
		return ParserCommandNotFound
	}
	parts := strings.Split(line, " ")
	if len(parts) <= 1 {
		// All commands, and their output, should
		// yield at least two parts.
		return ParserCommandNotFound
	}
	switch parts[0] {
	case "$":
		return parseCommand(parts[1:])
	case "dir":
		return ParserFoundDir
	default:
		return ParserFoundFile
	}
}

func parseCommand(commandParts []string) ParserAction {
	command, rest := h.Pop(commandParts)
	switch command {
	case "cd":
		switch rest[0] {
		case "/":
			return ParserCdRoot
		case "..":
			return ParserCdParent
		default:
			return ParserCdChild
		}
	case "ls":
		return ParserListCwd
	default:
		return ParserCommandNotFound
	}
}

// Statable items each define their own methods
// for identifying themselves.
type Statable interface {
	Name() string
	Size() int
	IsFile() bool
	IsDir() bool
}

// Files are meant to model a single file or "leaf" node.
// Implementes Statable.
type File struct {
	name string
	size int
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}

func (f File) IsFile() bool {
	return true
}

func (f File) IsDir() bool {
	return false
}

func (f File) AddTo(d *Directory) {
	d.children = append(d.children, &f)
}

func (f File) String() string {
	return fmt.Sprintf("[f] %s (%d)", f.name, f.size)
}

func NewFile(name string, size int) File {
	return File{name, size}
}

// Directories are meant to model containers for files and are a "branch"
// point.
// Implementes Statable.
type Directory struct {
	name     string
	children []Statable
	parent   *Directory
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Size() int {
	if len(d.children) == 0 {
		return 0
	}
	var sum int
	for _, f := range d.children {
		sum += f.Size()
	}
	return sum
}

func (d Directory) IsFile() bool {
	return false
}

func (d Directory) IsDir() bool {
	return true
}

func (d Directory) Parent() *Directory {
	return d.parent
}

func (d *Directory) AddParent(pd *Directory) {
	d.parent = pd
	// naive and surely a problem
	pd.children = append(pd.children, d)
}

func (d Directory) String() string {
	return fmt.Sprintf("[d] %s", d.name)
}

func NewDirectory(name string, children ...Statable) Directory {
	if len(children) == 0 {
		return Directory{name: name, children: make([]Statable, 0)}
	}
	return Directory{name: name, children: children}
}
