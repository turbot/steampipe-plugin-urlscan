package urlscan

import (
	"context"

	"github.com/m-mizutani/urlscan-go/urlscan"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanSearch(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_search",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listSearch,
			KeyColumns: plugin.SingleColumn("query"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the scan yielding this search result."},
			{Name: "page_domain", Type: proto.ColumnType_STRING, Transform: transform.FromField("Page.Domain"), Description: "Domain of the page."},
			{Name: "page_ip", Type: proto.ColumnType_IPADDR, Transform: transform.FromField("Page.IP"), Description: "IP address of the page."},
			{Name: "page_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Page.URL"), Description: "URL of the page."},
			{Name: "task_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Task.Time"), Description: "Time when the scan was run."},
			{Name: "task_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Task.URL"), Description: "URL of the task scan result."},
			{Name: "task_visibility", Type: proto.ColumnType_STRING, Transform: transform.FromField("Task.Visibility"), Description: "Visibility of the task scan."},
			// Other columns
			{Name: "page", Type: proto.ColumnType_JSON, Description: "Basic page information."},
			{Name: "query", Type: proto.ColumnType_STRING, Transform: transform.FromQual("query"), Description: "The search query."},
			{Name: "result", Type: proto.ColumnType_STRING, Description: "URL to the full scan result data."},
			{Name: "stats", Type: proto.ColumnType_JSON, Description: "Basic statistics for the page."},
			{Name: "task", Type: proto.ColumnType_JSON, Description: "Task information for the scan."},
		},
	}
}

func listSearch(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	query := quals["query"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_search.listSearch", "connection_error", err)
		return nil, err
	}
	// NOTE: Pagination is not yet implemented. The urlscan-go library uses an
	// "offset" parameter, but the actual search parameter is called
	// search_after. The expected URL is of a form
	// https://urlscan.io/api/v1/search?q=domain:github.io&size=1&search_after=1624767767102,422c855e-e916-4db4-af8e-1da2379793ad
	// passing the sort values of the last search result into the search_after param.
	params := urlscan.SearchArguments{
		Query: urlscan.String(query),
		Size:  urlscan.Uint64(10000),
	}
	results, err := conn.Search(params)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_search.listSearch", "query_error", err, "results", results)
		return nil, err
	}
	for _, i := range results.Results {
		d.StreamListItem(ctx, i)
	}
	return nil, nil
}
