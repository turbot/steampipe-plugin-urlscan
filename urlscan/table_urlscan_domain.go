package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_domain",
		Description: "Domains used by requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listDomain,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "The domain."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of requests to the domain."},
			// Other columns
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the domain resides."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests sent to this domain."},
			{Name: "initiators", Type: proto.ColumnType_JSON, Description: "Initiators of requests to this domain."},
			{Name: "ips", Type: proto.ColumnType_JSON, Description: "IP addresses for the domain."},
			{Name: "redirects", Type: proto.ColumnType_INT, Description: "Number of redirect responses received from this domain."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests sent to this domain."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_domain.listDomain", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_domain.listDomain", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.DomainStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
