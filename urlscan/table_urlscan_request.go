package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableURLScanRequest(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "urlscan_request",
		Description: "Requests configured for the account.",
		List: &plugin.ListConfig{
			Hydrate:    listRequest,
			KeyColumns: plugin.SingleColumn("scan"),
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "timestamp", Type: proto.ColumnType_DOUBLE, Transform: transform.FromField("Request.Timestamp"), Description: "Request timestamp received from urlscan. (Unknown format?)"},
			{Name: "method", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.Method"), Description: "HTTP method used for the request."},
			{Name: "url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.URL"), Description: "URL for the request."},
			{Name: "document_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Type"), Description: "Type of document requested."},
			{Name: "response_status", Type: proto.ColumnType_INT, Transform: transform.FromField("Response.Response.Status"), Description: "HTTP status code of the response, e.g. 200."},

			// Request columns
			{Name: "document_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.DocumentURL"), Description: "URL of the document the requests are building."},
			{Name: "frame_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.FrameID"), Description: "Frame ID in the request."},
			{Name: "has_post_data", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Request.Request.HasPostData"), Description: "True if the request includes post data."},
			{Name: "has_user_gesture", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Request.HasUserGesture"), Description: "True if the request has a user gesture."},
			{Name: "headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Request.Request.Headers"), Description: "Request headers."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.RequestID"), Description: "ID of the request."},
			{Name: "initial_priority", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.InitialPriority"), Description: "Initial priority for the request."},
			{Name: "initiator", Type: proto.ColumnType_JSON, Transform: transform.FromField("Request.Initiator"), Description: "Initiator (e.g. page) that made the request."},
			{Name: "loader_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.LoaderID"), Description: "ID of the loader for the request."},
			{Name: "mixed_content_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.MixedContentType"), Description: "Mixed content type information."},
			{Name: "post_data", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.PostData"), Description: "HTTP post data sent with the request."},
			{Name: "referrer_policy", Type: proto.ColumnType_STRING, Transform: transform.FromField("Request.Request.ReferrerPolicy"), Description: "Referrer policy for the request."},
			{Name: "wall_time", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Request.WallTime").Transform(transform.ToInt).Transform(transform.UnixToTimestamp), Description: "Clock time when the request was made."},

			// Response columns
			{Name: "response_abp", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Abp"), Description: "ABP information for the response."},
			{Name: "response_asn", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Asn"), Description: "ASN information for the IP address handling the request."},
			{Name: "response_data_length", Type: proto.ColumnType_INT, Transform: transform.FromField("Response.DataLength"), Description: "Response data length in bytes."},
			{Name: "response_encoded_data_length", Type: proto.ColumnType_INT, Transform: transform.FromField("Response.EncodedDataLength"), Description: "Response encoded data length in bytes."},
			{Name: "response_geolocation", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Geoip"), Description: "Geolocation information for the IP address handling the request."},
			{Name: "response_hash", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Hash"), Description: "Hash of the response."},
			{Name: "response_hash_matches", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Hashmatches"), Description: "Hash matches for the response."},
			{Name: "response_headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Response.Headers"), Description: "Response headers."},
			{Name: "response_mime_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Response.MimeType"), Description: "Mime type of the response."},
			{Name: "response_protocol", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Response.Protocol"), Description: "HTTP protocol for the response, e.g. h2."},
			{Name: "response_remote_ip_address", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Response.RemoteIPAddress"), Description: "IP address of the server responding to the request."},
			{Name: "response_remote_port", Type: proto.ColumnType_INT, Transform: transform.FromField("Response.Response.RemotePort"), Description: "Port for the server responding to the request."},
			{Name: "response_reverse_dns", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Rdns"), Description: "Reverse DNS information for the IP address handling the request."},
			{Name: "response_security_details", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Response.SecurityDetails"), Description: "Security details for the response."},
			{Name: "response_security_headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Response.SecurityHeaders"), Description: "Security headers for the response."},
			{Name: "response_security_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Response.SecurityState"), Description: "Security state for the response."},
			{Name: "response_size", Type: proto.ColumnType_INT, Transform: transform.FromField("Response.Size"), Description: "Size of the response."},
			{Name: "response_status_text", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Response.StatusText"), Description: "HTTP status text for the response."},
			{Name: "response_timing", Type: proto.ColumnType_JSON, Transform: transform.FromField("Response.Timing"), Description: "Timing data for the response."},
			{Name: "response_document_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Response.Type"), Description: "Type of document returned in the response."},

			// Context
			{Name: "scan", Type: proto.ColumnType_STRING, Transform: transform.FromQual("scan"), Description: "ID of the scan result."},
		},
	}
}

func listRequest(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	scan := quals["scan"].GetStringValue()
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("urlscan_request.listRequest", "connection_error", err)
		return nil, err
	}
	task := conn.ResultTask(scan)
	err = task.Get()
	if err != nil {
		if isNotFoundError(err) {
			return nil, nil
		}
		plugin.Logger(ctx).Error("urlscan_request.listRequest", "wait_error", err)
		return nil, err
	}
	for _, request := range task.Result.Data.Requests {
		d.StreamListItem(ctx, request)
	}
	return nil, nil
}
