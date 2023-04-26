package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/o98k-ok/lazy/v2/collection"
	"github.com/o98k-ok/vscode-remote-flow/command"
	"github.com/o98k-ok/vscode-remote-flow/config"
)

func main() {
	// init something
	cfg := config.NewConfig()

	cli := alfred.NewApp("vscode util toools")
	cli.Bind("get", func(s []string) {
		cmds := cfg.CmdConfig
		for _, key := range s {
			cmds = collection.SearchMap(cmds, key)
		}

		msg := alfred.NewItems()
		for k, c := range cmds {
			title := c.GetCommand().Name
			subtitle := fmt.Sprintf("已经使用%d次 [%s]", c.GetCommand().Count, k)

			item := alfred.NewItem(title, subtitle, k)
			item.Icon = &alfred.Icon{}
			item.WithIcon(fmt.Sprintf("./icons/%s.png", c.GetCommand().App))
			msg.Append(item)
		}

		sort.Slice(msg.Items, func(i, j int) bool {
			var left, right int
			fmt.Sscanf(msg.Items[i].SubTitle, "已经使用%d次", &left)
			fmt.Sscanf(msg.Items[j].SubTitle, "已经使用%d次", &right)
			return left > right
		})
		msg.Show()
	})
	cli.Bind("inc", func(s []string) {
		if len(s) <= 0 {
			alfred.InputErrItems("param size error").Show()
			return
		}

		cmd, ok := cfg.CmdConfig[s[0]]
		if !ok {
			alfred.EmptyItems().Show()
			return
		}

		cmd.IncCount()
		fmt.Println(cmd.GenURI())
		cfg.Save()
	})
	cli.Bind("add", func(s []string) {
		if len(s) < 4 {
			alfred.InputErrItems("input params size < 2").Show()
			return
		}

		key, name, app, cmdline := s[0], s[1], s[2], s[3]
		cmd := command.NewCommander(command.NewCommand(app, key, cmdline, cfg.ObsidianVault))
		cfg.CmdConfig[name] = cmd

		cfg.Save()
	})

	err := cli.Run(os.Args)
	if err != nil {
		alfred.ErrItems("run failed", err).Show()
		return
	}
}
