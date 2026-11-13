package urlscan

import (
	"context"
	"errors"
	"os"

	"github.com/m-mizutani/urlscan-go/urlscan"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*urlscan.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "urlscan"
	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return cachedData.(*urlscan.Client), nil
	}

	// Default to using env vars
	apiKey := os.Getenv("URLSCAN_API_KEY")

	// But prefer the config
	urlscanConfig := GetConfig(d.Connection)
	if urlscanConfig.APIKey != nil {
		apiKey = *urlscanConfig.APIKey
	}

	if apiKey == "" {
		// Credentials not set
		return nil, errors.New("api_key must be configured")
	}

	connItem := urlscan.NewClient(apiKey)
	conn := &connItem

	// Save to cache
	if err := d.ConnectionCache.Set(ctx, cacheKey, conn); err != nil {
		plugin.Logger(ctx).Warn("connect", "cache_set_error", err)
	}

	return conn, nil
}
