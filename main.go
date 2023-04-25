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
			title := c.Name
			subtitle := fmt.Sprintf("已经使用%d次 [%s]", c.Count, k)

			item := alfred.NewItem(title, subtitle, k)
			item.Icon = &alfred.Icon{}
			item.WithIcon(fmt.Sprintf("./icons/%s.png", c.App))
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

		// websocket, err := ws.NewCommandSocket(cfg.WsConfig)
		// if err != nil {
		// 	alfred.ErrItems("init websocket failed", err).Show()
		// 	return
		// }
		// defer websocket.Close()

		cmd, ok := cfg.CmdConfig[s[0]]
		if !ok {
			alfred.EmptyItems().Show()
			return
		}

		cmd.IncCount()
		// websocket.Do(cmd.Command)
		fmt.Println(fmt.Sprintf(cmd.Format, cmd.App, cmd.Command))
		cfg.Save()
	})
	cli.Bind("add", func(s []string) {
		if len(s) < 2 {
			alfred.InputErrItems("input params size < 2").Show()
			return
		}

		name, cmdline := s[0], s[1]
		cmd := command.NewCommand(name, cmdline)
		cfg.CmdConfig[name] = cmd

		cfg.Save()
	})

	err := cli.Run(os.Args)
	if err != nil {
		alfred.ErrItems("run failed", err).Show()
		return
	}
}
