package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableURLScanLink(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_link",
		Description: "Links in the HTML of the page.",
		List: &plugin.ListConfig{
			Hydrate:    listLink,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "href", Type: proto.ColumnType_STRING, Description: "HREF of the link."},
			{Name: "text", Type: proto.ColumnType_STRING, Description: "Text of the link."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listLink(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_link.listLink", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_link.listLink", "wait_error", err)
		return nil, err
	}
	for _, link := range task.Result.Data.Links {
		d.StreamListItem(ctx, link)
	}
	return nil, nil
}
