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

	ObsidianVault = "obsidian_vault"
)

var (
	DefaultConfigCommand = map[string]command.Commander{
		"open_terminal":    command.NewCommander(command.NewCommand(command.VSCODE, "open_terminal", "workbench.action.terminal.toggleTerminal", "")),
		"open_main_folder": command.NewCommander(command.NewCommand(command.VSCODE, "open_main_folder", "workbench.explorer.fileView.focus", "")),
		"extension_list":   command.NewCommander(command.NewCommand(command.VSCODE, "extension_list", "workbench.extensions.action.installExtensions", "")),
		"open_notes":       command.NewCommander(command.NewCommand(command.VSCODE, "open_notes", "workbench.view.extension.vscode-notes", "")),
		"list_projects":    command.NewCommander(command.NewCommand(command.VSCODE, "list_projects", "projectManager.listProjects", "")),
		"select_themes":    command.NewCommander(command.NewCommand(command.VSCODE, "select_themes", "workbench.action.selectTheme", "")),
		"open_shortcuts":   command.NewCommander(command.NewCommand(command.VSCODE, "open_shortcuts", "workbench.action.openGlobalKeybindings", "")),
		"open_settings":    command.NewCommander(command.NewCommand(command.VSCODE, "open_settings", "workbench.action.openSettings", "")),
		"command_window":   command.NewCommander(command.NewCommand(command.VSCODE, "command_window", "workbench.action.showCommands", "")),
		"view_markdown":    command.NewCommander(command.NewCommand(command.VSCODE, "view_markdown", "markdown.showPreview", "")),
		"close_right":      command.NewCommander(command.NewCommand(command.VSCODE, "close_right", "workbench.action.closeEditorsToTheRight", "")),
	}
	DefaultConfig = Config{
		DefaultWebSocket,
		DefaultJsonPath,
		command.OBSIDIAN,
		DefaultConfigCommand,
	}
)

type Config struct {
	WsConfig      string
	JsonFile      string
	ObsidianVault string
	CmdConfig     map[string]command.Commander
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

		vault, ok := envs[ObsidianVault]
		if ok && len(vault) != 0 {
			config.ObsidianVault = vault
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
	v.IncCount()
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
