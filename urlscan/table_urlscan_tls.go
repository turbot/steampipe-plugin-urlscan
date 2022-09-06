package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableURLScanTLS(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_tls",
		Description: "Transport Layer Security (TLS) status for requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listTLS,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "security_state", Type: proto.ColumnType_STRING, Description: "Security state for these requests."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of requests using this security."},
			// Other columns
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the tls resides."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests sent to this tls."},
			{Name: "ips", Type: proto.ColumnType_JSON, Description: "IP addresses for the tls."},
			{Name: "protocols", Type: proto.ColumnType_JSON, Description: "TLS protocols used by requests, including the count."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests sent to this tls."},
		},
	}
}

func listTLS(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_tls.listTLS", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_tls.listTLS", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.TLSStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
