package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableURLScanAsn(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_asn",
		Description: "Autonomous System Numbers (ASNs) used by requests during the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listAsn,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "asn", Type: proto.ColumnType_STRING, Description: "The ASN."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the ASN."},
			// Other columns
			{Name: "country", Type: proto.ColumnType_STRING, Description: "Country where the IP address was found."},
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Description: "Date when the ASN was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the ASN."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "IP address in the ASN."},
			{Name: "registrar", Type: proto.ColumnType_STRING, Description: "Registar for the ASN."},
			{Name: "route", Type: proto.ColumnType_CIDR, Description: "Route / CIDR for the ASN."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listAsn(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_asn.listAsn", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_asn.listAsn", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Meta.Processors.Asn.Data {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
