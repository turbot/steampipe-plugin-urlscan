package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanMetadata(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_metadata",
		Description: "Metadata details of the scan.",
		List: &plugin.ListConfig{
			Hydrate:    listMetadata,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},

			// Other columns
			{Name: "abp_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Abp.State"), Description: "ABP status for the scan metadata."},
			{Name: "abp_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Abp.Data"), Description: "ABP information for the scan metadata."},
			{Name: "asn_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Asn.State"), Description: "IP ASN annotation state for every IP contacted during page navigation."},
			{Name: "asn_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Asn.Data"), Description: "IP ASN annotation for every IP contacted during page navigation."},
			{Name: "download_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Download.State"), Description: "The state of the file download."},
			{Name: "download_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Download.Data"), Description: "List of files that were downloaded by the website, Keys: filename, filesize, receivedBytes, url, startedAt, state, mimeType, mimeDescription, sha256, finishedAt."},
			{Name: "geoip_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Geoip.State"), Description: "The state of GeoIP annotation for every IP contacted during page navigation."},
			{Name: "geoip_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Geoip.Data"), Description: "GeoIP annotation for every IP contacted during page navigation."},
			{Name: "rdns_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Rdns.State"), Description: "The state of DNS PTR records for every hostname contacted during page navigation."},
			{Name: "rdns_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Rdns.Data"), Description: "DNS PTR records for every hostname contacted during page navigation."},
			{Name: "wappa_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Wappa.State"), Description: "The state of Wappalyzer technology detection for fully loaded page."},
			{Name: "wappa_data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Wappa.Data"), Description: "Wappalyzer technology detection for fully loaded page."},
		},
	}
}

func listMetadata(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()

	// Empty check
	if scan == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_metadata.listMetadata", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_metadata.listMetadata", "wait_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, task.Result.Meta.Processors)
	return nil, nil
}
