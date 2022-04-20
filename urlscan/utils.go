package urlscan

import (
	"context"
	"errors"
	"os"

	"github.com/m-mizutani/urlscan-go/urlscan"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*urlscan.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "urlscan"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*urlscan.Client), nil
	}

	// Default to using env vars
	apiKey := os.Getenv("URLSCAN_API_KEY")

	// But prefer the config
	urlscanConfig := GetConfig(d.Connection)
	if &urlscanConfig != nil {
		if urlscanConfig.APIKey != nil {
			apiKey = *urlscanConfig.APIKey
		}
	}

	if apiKey == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}

	connItem := urlscan.NewClient(apiKey)
	conn := &connItem

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, conn)

	return conn, nil
}
