package command

import (
	"fmt"
	"strings"
)

const (
	DEFAULT_APPLE_FORMAT = "osascript://%s"
	APP_APPLE            = "apple"
	APPLE_CMD            = "osascript_cmd_"
)

type AppleCMD struct {
	format string
	script string
	attrs  map[string]string
}

func NewAppleCMD(commandID string, attrs map[string]string) *AppleCMD {
	return &AppleCMD{
		format: DEFAULT_APPLE_FORMAT,
		script: commandID,
		attrs:  attrs,
	}
}

func (oc *AppleCMD) GenURI() string {
	return fmt.Sprintf(oc.format, oc.script)
}

func (oc *AppleCMD) IconApp() string {
	return APP_APPLE
}

func (oc *AppleCMD) Filtered(keys []string) (string, string, bool) {
	for k, v := range oc.attrs {
		for _, query := range keys {
			lowk, lowquery := strings.ToLower(k), strings.ToLower(query)
			if !strings.Contains(lowk, lowquery) && !strings.Contains(v, lowquery) {
				return "", "", false
			}
		}
		return k, v, true
	}
	return "", "", false
}
