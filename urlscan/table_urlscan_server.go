package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanServer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_server",
		Description: "Servers used by requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listServer,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "server", Type: proto.ColumnType_STRING, Description: "The server."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of requests to the server."},
			// Other columns
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the server resides."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests sent to this server."},
			{Name: "ips", Type: proto.ColumnType_JSON, Description: "IP addresses for the server."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests sent to this server."},
		},
	}
}

func listServer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_server.listServer", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_server.listServer", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.ServerStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
