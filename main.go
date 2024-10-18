package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/mozillazg/go-pinyin"
	"github.com/o98k-ok/lazy/v2/alfred"
)

var FILE = "./config.json"

func readFile(file string) map[string]string {
	content, err := os.ReadFile(file)
	if err != nil {
		alfred.Log("raw %v", err.Error())
		return nil
	}

	var result map[string]string
	err = json.Unmarshal(content, &result)
	if err != nil {
		alfred.Log("raw %v", err.Error())
		return nil
	}
	return result
}

func readFileToMarkdown(file string) string {
	result := readFile(file)
	arr := []string{}
	for k, v := range result {
		arr = append(arr, fmt.Sprintf("* %s: %s", k, v))
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	response := map[string]string{
		"response": strings.Join(arr, "\n"),
	}
	data, err := json.Marshal(response)
	if err != nil {
		alfred.Log("raw %v", err.Error())
		return ""
	}
	return string(data)
}

func toPinyin(s string) string {
	return strings.Join(pinyin.LazyConvert(s, nil), "")
}

func main() {
	app := alfred.NewApp("command pro")
	app.Bind("edit", func(s []string) {
		content := os.Args[2]
		if len(content) == 0 {
			fmt.Println(readFileToMarkdown(FILE))
			return
		}

		_, err := url.Parse(content)
		if err == nil {
			fmt.Println(readFileToMarkdown(FILE))
			return
		}

		fields := strings.SplitN(content, ":", 2)
		if len(fields) != 2 {
			alfred.Log("edit %v", "json format invalid")
			return
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		if len(key) == 0 || len(value) == 0 {
			alfred.Log("edit %v", "key or value is empty")
			return
		}

		result := readFile(FILE)
		result[key] = value
		data, err := json.Marshal(result)
		if err != nil {
			alfred.Log("edit %v", "json format invalid")
			return
		}

		file, _ := os.OpenFile(FILE, os.O_WRONLY|os.O_TRUNC, 0o644)
		defer file.Close()
		file.Write(data)

		fmt.Println(readFileToMarkdown(FILE))
	})

	app.Bind("search", func(s []string) {
		all := readFile(FILE)
		filtered := map[string]string{}
		var content string
		if len(os.Args) > 2 {
			content = os.Args[2]
		}

		home, _ := os.Getwd()
		for k, v := range all {
			v = strings.ReplaceAll(v, "{{HOME}}", home)
			if strings.Contains(v, content) ||
				strings.Contains(toPinyin(k), content) ||
				strings.Contains(k, content) {
				filtered[k] = v
			}
		}

		// add sort by app
		items := alfred.NewItems()
		items2 := alfred.NewItems()
		frontmostAppName, _ := getFrontmostAppName()
		for k, v := range filtered {
			appName := getAppNameFromLink(v)
			item := alfred.NewItem(k, v, v).WithIcon("./icons/" + appName + ".png")
			if appName == frontmostAppName {
				items.Append(item)
			} else {
				items2.Append(item)
			}
		}

		items.Items = append(items.Items, items2.Items...)
		items.Show()
	})

	app.Run(os.Args)
}

// getAppNameFromLink 获取链接的应用名
// 支持的链接格式：
// bash://xxx?app=iTerm2
// scpt://xxx?app=iTerm2
// 其他格式直接返回 scheme
func getAppNameFromLink(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	switch u.Scheme {
	case "bash", "scpt":
		return u.Query().Get("app")
	default:
		return u.Scheme
	}
}

// getFrontmostAppName 获取前台应用名
func getFrontmostAppName() (string, error) {
	cmd := exec.Command("osascript", "-e", `
		tell application "System Events"
			set frontmostProcess to first process where it is frontmost
			set appName to name of frontmostProcess
			return appName
		end tell
	`)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	app := strings.TrimSpace(string(output))

	switch app {
	case "Google Chrome":
		return "chrome", nil
	default:
		return strings.ToLower(app), nil
	}
}
