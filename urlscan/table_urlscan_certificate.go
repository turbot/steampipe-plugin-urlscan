package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableURLScanCertificate(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_certificate",
		Description: "Certificates returned by requests for the page.",
		List: &plugin.ListConfig{
			Hydrate:    listCertificate,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "subject_name", Type: proto.ColumnType_STRING, Description: "Subject (domain name) of the certificate."},
			{Name: "issuer", Type: proto.ColumnType_STRING, Description: "Issuer of the certificate."},
			{Name: "valid_from", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ValidFrom").Transform(transform.UnixToTimestamp), Description: "Time the certificate is valid from."},
			{Name: "valid_to", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ValidTo").Transform(transform.UnixToTimestamp), Description: "Time the certificate is valid to."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_certificate.listCertificate", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_certificate.listCertificate", "wait_error", err)
		return nil, err
	}
	for _, certificate := range task.Result.Lists.Certificates {
		d.StreamListItem(ctx, certificate)
	}
	return nil, nil
}
