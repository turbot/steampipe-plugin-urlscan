package urlscan

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-urlscan",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"urlscan_app":           tableURLScanApp(ctx),
			"urlscan_asn":           tableURLScanAsn(ctx),
			"urlscan_certificate":   tableURLScanCertificate(ctx),
			"urlscan_console":       tableURLScanConsole(ctx),
			"urlscan_cookie":        tableURLScanCookie(ctx),
			"urlscan_domain":        tableURLScanDomain(ctx),
			"urlscan_global":        tableURLScanGlobal(ctx),
			"urlscan_ip":            tableURLScanIP(ctx),
			"urlscan_link":          tableURLScanLink(ctx),
			"urlscan_metadata":      tableURLScanMetadata(ctx),
			"urlscan_protocol":      tableURLScanProtocol(ctx),
			"urlscan_request":       tableURLScanRequest(ctx),
			"urlscan_document_type": tableURLScanDocumentType(ctx),
			"urlscan_search":        tableURLScanSearch(ctx),
			"urlscan_server":        tableURLScanServer(ctx),
			"urlscan_tls":           tableURLScanTLS(ctx),
			"urlscan_task":          tableURLScanTask(ctx),
		},
	}
	return p
}
