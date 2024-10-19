package task

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/o98k-ok/command/internel/config"
	"github.com/o98k-ok/command/internel/textview"
	"github.com/o98k-ok/command/internel/toolkit"
)

func EditTask(s []string) {
	configFile := config.GetConfigFile()
	var content string
	if len(s) > 0 {
		content = s[0]
	}
	_, err := url.Parse(content)
	if err == nil || len(content) == 0 {
		result, err := textview.FileToTextViewMarkdown(configFile)
		toolkit.ShowByAlfred("EditTask", result, err)
		return
	}

	fields := strings.SplitN(content, ":", 2)
	if len(fields) != 2 {
		toolkit.ShowByAlfred("EditTask", "", errors.New("json format invalid"))
		return
	}
	key := strings.TrimSpace(fields[0])
	value := strings.TrimSpace(fields[1])
	if len(key) == 0 || len(value) == 0 {
		toolkit.ShowByAlfred("EditTask", "", errors.New("key or value is empty"))
		return
	}

	result, err := textview.ReadJSON(configFile)
	if err != nil {
		toolkit.ShowByAlfred("EditTask", "", err)
		return
	}
	result[key] = value

	var file *os.File
	if file, err = os.OpenFile(configFile, os.O_WRONLY|os.O_TRUNC, 0o644); err != nil {
		toolkit.ShowByAlfred("EditTask", "", err)
		return
	}
	defer file.Close()
	if err = json.NewEncoder(file).Encode(result); err != nil {
		toolkit.ShowByAlfred("EditTask", "", err)
		return
	}

	output, err := textview.FileToTextViewMarkdown(configFile)
	toolkit.ShowByAlfred("EditTask", output, err)
}
