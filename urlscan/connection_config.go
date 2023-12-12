package urlscan

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type urlscanConfig struct {
	APIKey *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &urlscanConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) urlscanConfig {
	if connection == nil || connection.Config == nil {
		return urlscanConfig{}
	}
	config, _ := connection.Config.(urlscanConfig)
	return config
}
