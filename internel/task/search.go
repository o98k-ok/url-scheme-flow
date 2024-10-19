package task

import (
	"os"
	"strings"

	"github.com/o98k-ok/command/internel/config"
	"github.com/o98k-ok/command/internel/textview"
	"github.com/o98k-ok/command/internel/toolkit"
	"github.com/o98k-ok/lazy/v2/alfred"
)

func SearchTask(s []string) {
	configFile := config.GetConfigFile()
	var query string
	if len(s) > 0 {
		query = s[0]
	}

	all, err := textview.ReadJSON(configFile)
	if err != nil {
		toolkit.ShowByAlfred("SearchTask", "", err)
		return
	}

	filtered := map[string]string{}
	home, _ := os.Getwd()
	for k, v := range all {
		v = strings.ReplaceAll(v, "{{HOME}}", home)
		if toolkit.PinyinMatchQuery(k, query) ||
			strings.Contains(v, query) {
			filtered[k] = v
		}
	}

	// add sort by app
	items := alfred.NewItems()
	items2 := alfred.NewItems()
	frontmostAppName, _ := toolkit.GetFrontmostAppName()
	for k, v := range filtered {
		appName := toolkit.GetAppNameFromLink(v)
		item := alfred.NewItem(k, v, v).WithIcon("./icons/" + appName + ".png")
		if appName == frontmostAppName {
			items.Append(item)
		} else {
			items2.Append(item)
		}
	}
	items.Items = append(items.Items, items2.Items...)
	items.Show()
}
