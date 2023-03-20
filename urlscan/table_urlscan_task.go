package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableURLScanTask(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_task",
		Description: "Details of the task that ran the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listTask,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "dom_url", Type: proto.ColumnType_STRING, Description: "DOM URL for the task."},
			{Name: "method", Type: proto.ColumnType_STRING, Description: "Method for the task, e.g. automatic."},
			// Omitted, only option is the user agent
			// {Name: "options", Type: proto.ColumnType_JSON, Description: "Options for the task."},
			{Name: "report_url", Type: proto.ColumnType_STRING, Description: "URL where the report is available."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "screenshot_url", Type: proto.ColumnType_STRING, Description: "URL for a screenshot of the page."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "Source that triggered the task."},
			{Name: "time", Type: proto.ColumnType_STRING, Description: "Time when the task was run."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the task."},
			{Name: "user_agent", Type: proto.ColumnType_STRING, Description: "User agent used for the task."},
			{Name: "uuid", Type: proto.ColumnType_STRING, Description: "UUID of the task."},
			{Name: "visibility", Type: proto.ColumnType_STRING, Description: "Visibility of the task, e.g. public."},
		},
	}
}

func listTask(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_task.listTask", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_task.listTask", "wait_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, task.Result.Task)
	return nil, nil
}
