package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableURLScanDocumentType(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_document_type",
		Description: "Document types returned for the page.",
		List: &plugin.ListConfig{
			Hydrate:    listDocumentType,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The document type, e.g. Script, Image."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "Number of requests for the document type."},
			// Other columns
			{Name: "countries", Type: proto.ColumnType_JSON, Description: "Countries where the document type was found."},
			{Name: "encoded_size", Type: proto.ColumnType_INT, Description: "Total encoded size of all requests for the document type."},
			{Name: "ips", Type: proto.ColumnType_JSON, Description: "IP addresses for the document type."},
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
			{Name: "size", Type: proto.ColumnType_INT, Description: "Total size of all requests for document type."},
			{Name: "compression", Type: proto.ColumnType_DOUBLE, Description: "Compression ratio for the document type."},
			{Name: "percentage", Type: proto.ColumnType_INT, Description: "Percentage of requests returning this document type."},
		},
	}
}

func listDocumentType(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_resource.listDocumentType", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_resource.listDocumentType", "wait_error", err)
		return nil, err
	}
	for _, item := range task.Result.Stats.ResourceStats {
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
