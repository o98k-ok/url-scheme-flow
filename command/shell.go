package command

import (
	"fmt"
	"strings"
)

const (
	DEFAULT_SHELL_FORMAT = "bash://%s"
	APP_SHELL            = "shell"
	SHELL_SCRIPT_CMD     = "shell_cmd_"
)

type ShellCMD struct {
	format string
	script string
	attrs  map[string]string
}

func NewShellCMD(commandID string, attrs map[string]string) *ShellCMD {
	return &ShellCMD{
		format: DEFAULT_SHELL_FORMAT,
		script: commandID,
		attrs:  attrs,
	}
}

func (oc *ShellCMD) GenURI() string {
	return fmt.Sprintf(oc.format, oc.script)
}

func (oc *ShellCMD) IconApp() string {
	return APP_SHELL
}

func (oc *ShellCMD) Filtered(keys []string) (string, string, bool) {
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

func (oc *ShellCMD) SetArgs(v string) {
}

func (oc *ShellCMD) Title() (string, string) {
	return oc.attrs[oc.script], oc.script
}
