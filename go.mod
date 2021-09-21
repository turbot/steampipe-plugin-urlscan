module github.com/turbot/steampipe-plugin-urlscan

go 1.16

replace github.com/m-mizutani/urlscan-go => github.com/e-gineer/urlscan-go v1.0.1-0.20210701205010-6cf8288d0d10

require (
	github.com/m-mizutani/urlscan-go v1.0.0
	github.com/turbot/steampipe-plugin-sdk v1.6.1
)
