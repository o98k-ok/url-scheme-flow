package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/o98k-ok/vscode-remote-flow/command"
)

const (
	obsidian_vault = "global_obsidian_vault"
	default_option = "global_default_choices"
	orderID        = "order="
)

func appendCommand(commands []command.Commander, key string, env string, obsidianVault string) []command.Commander {
	switch {
	case strings.HasPrefix(key, command.OB_CMD) && len(obsidianVault) != 0:
		commands = append(commands, command.NewObsidianCMD(env, obsidianVault, map[string]string{env: key}))
	case strings.HasPrefix(key, command.VS_CMD):
		commands = append(commands, command.NewVscodeCMD(env, map[string]string{env: key}))
	case strings.HasPrefix(key, command.LARK_CHAT_CMD):
		commands = append(commands, command.NewLarkChatCMD(env, map[string]string{env: key}))
	case strings.HasPrefix(key, command.SHELL_SCRIPT_CMD):
		commands = append(commands, command.NewShellCMD(env, map[string]string{env: key}))
	case strings.HasPrefix(key, command.APPLE_CMD):
		commands = append(commands, command.NewAppleCMD(env, map[string]string{env: key}))
	default:
	}
	return commands
}

func main() {
	envs, err := alfred.FlowVariables()
	if err != nil {
		alfred.InputErrItems("read env failed " + err.Error()).Show()
		return
	}

	// try get obsidian vault
	obsidianVault := envs[obsidian_vault]
	var commands []command.Commander
	for key, env := range envs {
		commands = appendCommand(commands, key, env, obsidianVault)
	}

	var optionsCommands []command.Commander
	choices := envs[default_option]
	if len(choices) != 0 {
		for _, c := range strutil.SplitEx(choices, ";", true) {
			key, env := c, envs[c]
			optionsCommands = appendCommand(optionsCommands, key, env, obsidianVault)
		}
	}

	cli := alfred.NewApp("vscode util toools")
	cli.Bind("get", func(query []string) {
		orderParams, params := slice.GroupBy(query, func(_ int, item string) bool { return strings.HasPrefix(item, orderID) })

		msg := alfred.NewItems()
		for _, cmd := range commands {
			key, name, ok := cmd.Filtered(params)
			if !ok {
				continue
			}

			item := alfred.NewItem(name, fmt.Sprintf("✌️✌️ %s", key), cmd.GenURI())
			item.Icon = &alfred.Icon{}
			item.WithIcon(fmt.Sprintf("./icons/%s.png", cmd.IconApp()))
			msg.Append(item)
		}

		// support order options
		if len(orderParams) != 0 {
			order, _ := strings.CutPrefix(orderParams[0], orderID)
			slice.SortBy(msg.Items, func(l *alfred.Item, r *alfred.Item) bool {
				return !strings.Contains(strings.ToLower(r.Title), order)
			})
		}

		// support bakeup options
		if len(params) != 0 {
			for _, cmd := range optionsCommands {
				cmd.SetArgs(params[0])
				title, _ := cmd.Title()
				item := alfred.NewItem(title, fmt.Sprintf("✌️✌️ %s", params[0]), cmd.GenURI())
				item.Icon = &alfred.Icon{}
				item.WithIcon(fmt.Sprintf("./icons/%s.png", cmd.IconApp()))
				msg.Append(item)
			}
		}
		msg.Show()
	})
	if err := cli.Run(os.Args); err != nil {
		alfred.ErrItems("run failed", err).Show()
		return
	}
}
