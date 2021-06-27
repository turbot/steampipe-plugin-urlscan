package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanGlobal(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_global",
		Description: "Global variables set in Javascript for the page.",
		List: &plugin.ListConfig{
			Hydrate:    listGlobal,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "property", Type: proto.ColumnType_STRING, Transform: transform.FromField("Prop"), Description: "Property name of the global variable."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the global variable."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listGlobal(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_global.listGlobal", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Wait()
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_global.listGlobal", "wait_error", err)
		return nil, err
	}
	for _, g := range task.Result.Data.Globals {
		d.StreamListItem(ctx, g)
	}
	return nil, nil
}
