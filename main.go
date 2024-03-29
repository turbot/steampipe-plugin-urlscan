package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-urlscan/urlscan"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: urlscan.Plugin})
}
