package toolkit

import (
	"net/url"
	"os/exec"
	"strings"
)

// GetAppNameFromLink 获取链接的应用名
// 支持的链接格式：
// bash://xxx?app=iTerm2
// scpt://xxx?app=iTerm2
// 其他格式直接返回 scheme
func GetAppNameFromLink(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		return ""
	}

	switch u.Scheme {
	case "bash", "scpt":
		return u.Query().Get("app")
	case "lark":
		return "feishu"
	default:
		return u.Scheme
	}
}

// GetFrontmostAppName 获取前台应用名
func GetFrontmostAppName() (string, error) {
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
