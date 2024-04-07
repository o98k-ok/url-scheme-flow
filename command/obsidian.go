package command

import (
	"fmt"
	"strings"
)

const (
	DEFAULT_OBSIDIAN_FORMAT = "%s://advanced-uri?vault=%s&commandid=%s"
	APP_OBSIDIAN            = "obsidian"
	DEFAULT_OBSIDIAN_VAULT  = "obsidian"
	OB_CMD                  = "obsidian_cmd_"
)

type ObidianCMD struct {
	format    string
	vault     string
	commandID string

	attrs map[string]string
}

func NewObsidianCMD(commandID, vault string, attrs map[string]string) *ObidianCMD {
	if len(vault) == 0 {
		vault = DEFAULT_OBSIDIAN_VAULT
	}

	return &ObidianCMD{
		format:    DEFAULT_OBSIDIAN_FORMAT,
		vault:     vault,
		commandID: commandID,
		attrs:     attrs,
	}
}

func (oc *ObidianCMD) GenURI() string {
	return fmt.Sprintf(oc.format, APP_OBSIDIAN, oc.vault, oc.commandID)
}

func (oc *ObidianCMD) IconApp() string {
	return APP_OBSIDIAN
}

func (oc *ObidianCMD) Filtered(keys []string) (string, string, bool) {
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
