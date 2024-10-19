package config

import (
	"github.com/o98k-ok/lazy/v2/alfred"
)

const CONFIG_FILE_KEY = "CONFIG_FILE"

func GetConfigFile() string {
	return alfred.GetEnvAsString(CONFIG_FILE_KEY, "config.json")
}
