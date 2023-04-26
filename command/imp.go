package command

import (
	"encoding/json"
	"fmt"
	"io"
)

type VscodeCMD struct {
	Command
}

func (vc *VscodeCMD) GenURI() string {
	return fmt.Sprintf(vc.Format, vc.App, vc.Command.Command)
}

type ObidianCMD struct {
	Command
}

func (oc *ObidianCMD) GenURI() string {
	fmt.Println(*oc)
	return fmt.Sprintf(oc.Format, oc.App, oc.Desc, oc.Command.Command)
}

type DefaultCMD struct {
	Command
}

func (dc *DefaultCMD) GenURI() string {
	return dc.Command.Command
}

func NewCommand(app, name, cmd, vault string) *Command {
	return &Command{
		App:     app,
		Name:    name,
		Command: cmd,
		Desc:    vault,
	}
}

func NewCommander(inner *Command) Commander {
	switch {
	case inner.App == VSCODE:
		inner.Format = DEFAULT_VSCODE_FORMAT
		return &VscodeCMD{*inner}
	case inner.App == OBSIDIAN:
		inner.Format = DEFAULT_OBSIDIAN_FORMAT
		return &ObidianCMD{*inner}
	default:
		inner.Format = ""
		return &DefaultCMD{*inner}
	}
}

func GetCommandsFrom(reader io.Reader) map[string]Commander {
	var config map[string]*Command
	err := json.NewDecoder(reader).Decode(&config)
	if err != nil {
		return map[string]Commander{}
	}

	result := make(map[string]Commander)
	for k, v := range config {
		result[k] = NewCommander(v)
	}
	return result
}
