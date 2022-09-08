package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableURLScanApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_app",
		Description: "Apps used by the page.",
		List: &plugin.ListConfig{
			Hydrate:    listApp,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "app", Type: proto.ColumnType_STRING, Description: "The app that was detected."},
			{Name: "website", Type: proto.ColumnType_STRING, Description: "Website of the app."},
			// Other columns
			{Name: "categories", Type: proto.ColumnType_JSON, Description: "Categories for the app."},
			{Name: "confidence", Type: proto.ColumnType_JSON, Description: "Confidence level in detection of the app."},
			{Name: "confidence_total", Type: proto.ColumnType_INT, Description: "Total confidence in the app detection as a perceentage."},
			{Name: "icon", Type: proto.ColumnType_STRING, Description: "Icon for the app."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_app.listApp", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_app.listApp", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Meta.Processors.Wappa.Data {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
