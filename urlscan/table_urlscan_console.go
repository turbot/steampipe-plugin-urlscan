package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableURLScanConsole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_console",
		Description: "Console messages returned while executing the page.",
		List: &plugin.ListConfig{
			Hydrate:    listConsole,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "rank", Type: proto.ColumnType_INT, Description: "Rank of the console message, used to maintain order."},
			{Name: "level", Type: proto.ColumnType_STRING, Transform: transform.FromField("Message.Level"), Description: "Level of the console message, e.g. log, warning."},
			{Name: "text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Message.Text"), Description: "Text of the console message."},
			// Other columns
			{Name: "column", Type: proto.ColumnType_INT, Transform: transform.FromField("Message.Column"), Description: "Column of code on the line that generated the console message."},
			{Name: "line", Type: proto.ColumnType_INT, Transform: transform.FromField("Message.Line"), Description: "Line of code that generated the console message."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "source", Type: proto.ColumnType_STRING, Transform: transform.FromField("Message.Source"), Description: "Source of the console message."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Message.URL"), Description: "URL of file that generated the console message."},
		},
	}
}

type consoleItem struct {
	Rank    int            `json:"rank"`
	Message consoleMessage `json:"message"`
}

type consoleMessage struct {
	Source string `json:"source"`
	Level  string `json:"level"`
	Text   string `json:"text"`
	URL    string `json:"url"`
	Line   int    `json:"line"`
	Column int    `json:"column"`
}

func listConsole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_console.listConsole", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_console.listConsole", "wait_error", err)
		return nil, err
	}
	for i, ci := range task.Result.Data.Console {
		c := ci.(map[string]interface{})
		m := c["message"].(map[string]interface{})
		cm := consoleMessage{
			Source: m["source"].(string),
			Level:  m["level"].(string),
			Text:   m["text"].(string),
			URL:    m["url"].(string),
			Line:   int(m["line"].(float64)),
			Column: int(m["column"].(float64)),
		}
		d.StreamListItem(ctx, consoleItem{Rank: i + 1, Message: cm})
	}
	return nil, nil
}
