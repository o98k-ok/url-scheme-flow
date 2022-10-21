package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/o98k-ok/lazy/v2/alfred"
	"github.com/o98k-ok/lazy/v2/collection"
	"github.com/o98k-ok/vscode-remote-flow/config"
	"github.com/o98k-ok/vscode-remote-flow/ws"
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
		for _, c := range cmds {
			msg.Append(alfred.NewItem(c.Name, fmt.Sprintf("has been used %d", c.Count), c.Name))
		}

		sort.Slice(msg.Items, func(i, j int) bool {
			var left, right int
			fmt.Sscanf(msg.Items[i].SubTitle, "has been used %d", &left)
			fmt.Sscanf(msg.Items[j].SubTitle, "has been used %d", &right)
			return left > right
		})
		msg.Show()
	})
	cli.Bind("inc", func(s []string) {
		if len(s) <= 0 {
			alfred.InputErrItems("param size error")
			return
		}

		websocket, err := ws.NewCommandSocket(cfg.WsConfig)
		if err != nil {
			alfred.ErrItems("init websocket failed", err).Show()
			return
		}
		defer websocket.Close()

		cmd, ok := cfg.CmdConfig[s[0]]
		if !ok {
			alfred.EmptyItems().Show()
			return
		}

		cmd.IncCount()
		websocket.Do(cmd.Command)
		cfg.Save()
	})

	err := cli.Run(os.Args)
	if err != nil {
		alfred.ErrItems("run failed", err).Show()
		return
	}
}
