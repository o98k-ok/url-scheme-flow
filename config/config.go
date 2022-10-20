package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/o98k-ok/vscode-remote-flow/command"
)

const (
	WSConfig         = "websocket_url"
	DefaultWebSocket = "ws://127.0.0.1:3710"

	CMDJsonPath     = "command_path"
	DefaultJsonPath = "config.json"
)

var (
	DefaultConfigCommand = map[string]*command.Command{
		"open_terminal":    command.NewCommand("open_terminal", "workbench.action.terminal.toggleTerminal"),
		"open_main_folder": command.NewCommand("open_main_folder", "workbench.explorer.fileView.focus"),
		"extension_list":   command.NewCommand("extension_list", "workbench.extensions.action.installExtensions"),
		"open_notes":       command.NewCommand("open_notes", "workbench.view.extension.vscode-notes"),
		"list_projects":    command.NewCommand("list_projects", "projectManager.listProjects"),
		"select_themes":    command.NewCommand("select_themes", "workbench.action.selectTheme"),
		"open_shortcuts":   command.NewCommand("open_shortcuts", "workbench.action.openGlobalKeybindings"),
		"open_settings":    command.NewCommand("open_settings", "workbench.action.openSettings"),
		"command_window":   command.NewCommand("command_window", "workbench.action.showCommands"),
		"view_markdown":    command.NewCommand("view_markdown", "markdown.showPreview"),
		"close_right":      command.NewCommand("close_right", "workbench.action.closeEditorsToTheRight"),
	}
	DefaultConfig = Config{
		DefaultWebSocket,
		DefaultJsonPath,
		DefaultConfigCommand,
	}
)

type Config struct {
	WsConfig  string
	JsonFile  string
	CmdConfig map[string]*command.Command
}

func NewConfig() *Config {
	var config = DefaultConfig
	envs, err := alfred.FlowVariables()
	if err == nil {
		url, ok := envs[WSConfig]
		if ok && len(url) != 0 {
			config.WsConfig = url
		}

		file, ok := envs[CMDJsonPath]
		if ok && len(file) != 0 {
			config.JsonFile = file
		}
	}

	f, err := os.Open(config.JsonFile)
	if err != nil {
		return &config
	}
	defer f.Close()

	cmd := command.GetCommandsFrom(f)
	if len(cmd) != 0 {
		config.CmdConfig = cmd
	}

	return &config
}

func (c *Config) IncCount(name string) {
	v, ok := c.CmdConfig[name]
	if !ok {
		return
	}
	v.Count++
}

func (c *Config) Save() error {
	v, err := json.Marshal(c.CmdConfig)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.JsonFile, v, 0644)
}

func (c *Config) String() string {
	var res strings.Builder
	for k, v := range c.CmdConfig {
		res.WriteString(fmt.Sprintf("%s:%v", k, v))
	}
	return res.String()
}
