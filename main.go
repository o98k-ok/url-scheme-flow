package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/o98k-ok/vscode-remote-flow/command"
)

const (
	obsidian_vault = "global_obsidian_vault"
)

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
	}

	cli := alfred.NewApp("vscode util toools")
	cli.Bind("get", func(s []string) {
		msg := alfred.NewItems()

		for _, cmd := range commands {
			key, name, ok := cmd.Filtered(s)
			if !ok {
				continue
			}

			item := alfred.NewItem(name, fmt.Sprintf("✌️✌️ %s", key), cmd.GenURI())
			item.Icon = &alfred.Icon{}
			item.WithIcon(fmt.Sprintf("./icons/%s.png", cmd.IconApp()))
			msg.Append(item)
		}

		msg.Show()
	})
	if err := cli.Run(os.Args); err != nil {
		alfred.ErrItems("run failed", err).Show()
		return
	}
}
