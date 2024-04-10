package command

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	DEFAULT_VSCODE_FORMAT = "%s://ionutvmi.vscode-commands-executor/runCommands?data=[%s]"
	APP_VSCODE            = "vscode"
	VS_CMD                = "vscode_cmd_"
)

type VscodeCMD struct {
	format    string
	commandID string
	args      string

	attrs map[string]string
}

type command struct {
	ID   string `json:"id"`
	Args string `json:"args,omitempty"`
}

func NewVscodeCMD(commandID string, attrs map[string]string) *VscodeCMD {
	return &VscodeCMD{
		format:    DEFAULT_VSCODE_FORMAT,
		commandID: commandID,

		attrs: attrs,
	}
}

func (vc *VscodeCMD) GenURI() string {
	d, _ := json.Marshal(command{ID: vc.commandID, Args: vc.args})
	return fmt.Sprintf(vc.format, APP_VSCODE, string(d))
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

func (vc *VscodeCMD) SetArgs(v string) {
	vc.args = v
}

func (oc *VscodeCMD) Title() (string, string) {
	return oc.attrs[oc.commandID], oc.commandID
}
