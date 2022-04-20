package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableURLScanIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_ip",
		Description: "IPs used by requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listIP,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "The IP address."},
			{Name: "count", Type: proto.ColumnType_INT, Transform: transform.FromField("Requests"), Description: "Number of requests to the IP."},
			// Other columns
			{Name: "asn", Type: proto.ColumnType_JSON, Description: "ASN data for this IP."},
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the IP resides."},
			{Name: "dns", Type: proto.ColumnType_JSON, Description: "DNS data for this IP."},
			{Name: "domains", Type: proto.ColumnType_JSON, Description: "List of domains used with this IP."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests sent to this IP."},
			{Name: "geolocation", Type: proto.ColumnType_JSON, Transform: transform.FromField("Geoip"), Description: "Geolocation data for this IP."},
			{Name: "ipv6", Type: proto.ColumnType_BOOL, Description: "True if this is an IPv6 address."},
			{Name: "redirects", Type: proto.ColumnType_INT, Description: "Number of redirect responses received from this IP."},
			{Name: "reverse_dns", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rdns.Ptr"), Description: "Reverse DNS lookup for the IP."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests sent to this IP."},
		},
	}
}

func listIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_ip.listIP", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_ip.listIP", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.IPStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
