package command

import (
	"encoding/json"
	"io"
)

type Command struct {
	Name    string
	Desc    string
	Command string
	Count   uint64
	Direct  bool
	App     string
	Format  string
}

func NewCommand(name, cmd string) *Command {
	return &Command{
		Name:    name,
		Command: cmd,
	}
}

func GetCommandsFrom(reader io.Reader) map[string]*Command {
	var config map[string]*Command
	err := json.NewDecoder(reader).Decode(&config)
	if err != nil {
		return map[string]*Command{}
	}

	return config
}

func (c *Command) IncCount() {
	c.Count++
}
