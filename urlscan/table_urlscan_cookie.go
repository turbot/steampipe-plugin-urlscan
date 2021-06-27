package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanCookie(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_cookie",
		Description: "Cookies set by the page for the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listCookie,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the cookie."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Value of the cookie."},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain of the cookie."},
			{Name: "expires", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Expires").Transform(transform.UnixToTimestamp), Description: "Time when the cookie expires."},
			{Name: "http_only", Type: proto.ColumnType_BOOL, Transform: transform.FromField("HTTPOnly"), Description: "True if the cookie can only be accessed via HTTP requests."},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "Path of the cookie."},
			{Name: "secure", Type: proto.ColumnType_BOOL, Description: "True if the cookie is only available over secure connections."},
			{Name: "session", Type: proto.ColumnType_BOOL, Description: "True if the cookie is only valid for the current session."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Size of the cookie."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listCookie(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_cookie.listCookie", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Wait()
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_cookie.listCookie", "wait_error", err)
		return nil, err
	}
	for _, cookie := range task.Result.Data.Cookies {
		d.StreamListItem(ctx, cookie)
	}
	return nil, nil
}
