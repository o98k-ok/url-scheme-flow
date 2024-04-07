package command

import (
	"fmt"
	"strings"
)

const (
	DEFAULT_VSCODE_FORMAT = "%s://ionutvmi.vscode-commands-executor/runCommands?data=[{\"id\": \"%s\"}]"
	APP_VSCODE            = "vscode"
	VS_CMD                = "vscode_cmd_"
)

type VscodeCMD struct {
	format    string
	commandID string

	attrs map[string]string
}

func NewVscodeCMD(commandID string, attrs map[string]string) *VscodeCMD {
	return &VscodeCMD{
		format:    DEFAULT_VSCODE_FORMAT,
		commandID: commandID,

		attrs: attrs,
	}
}

func (vc *VscodeCMD) GenURI() string {
	return fmt.Sprintf(vc.format, APP_VSCODE, vc.commandID)
}

func (vc *VscodeCMD) IconApp() string {
	return APP_VSCODE
}

func (vc *VscodeCMD) Filtered(keys []string) (string, string, bool) {
	for k, v := range vc.attrs {
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
