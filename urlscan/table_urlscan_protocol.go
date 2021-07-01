package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanProtocol(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_protocol",
		Description: "Protocols used by requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listProtocol,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "protocol", Type: proto.ColumnType_STRING, Description: "The protocol."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of requests to the protocol."},
			// Other columns
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the protocol resides."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests sent to this protocol."},
			{Name: "ips", Type: proto.ColumnType_JSON, Description: "IP addresses for the protocol."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests sent to this protocol."},
		},
	}
}

func listProtocol(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_protocol.listProtocol", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_protocol.listProtocol", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.ProtocolStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
