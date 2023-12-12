## v0.6.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#33](https://github.com/turbot/steampipe-plugin-urlscan/pull/33))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#33](https://github.com/turbot/steampipe-plugin-urlscan/pull/33))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-urlscan/blob/main/docs/LICENSE). ([#33](https://github.com/turbot/steampipe-plugin-urlscan/pull/33))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#32](https://github.com/turbot/steampipe-plugin-urlscan/pull/32))

## v0.5.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#25](https://github.com/turbot/steampipe-plugin-urlscan/pull/25))

## v0.5.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#23](https://github.com/turbot/steampipe-plugin-urlscan/pull/23))
- Recompiled plugin with Go version `1.21`. ([#23](https://github.com/turbot/steampipe-plugin-urlscan/pull/23))

## v0.4.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#19](https://github.com/turbot/steampipe-plugin-urlscan/pull/19))

## v0.3.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#17](https://github.com/turbot/steampipe-plugin-urlscan/pull/17))
- Recompiled plugin with Go version `1.19`. ([#17](https://github.com/turbot/steampipe-plugin-urlscan/pull/17))

## v0.2.1 [2022-05-24]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#13](https://github.com/turbot/steampipe-plugin-urlscan/pull/13))

## v0.2.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#11](https://github.com/turbot/steampipe-plugin-urlscan/pull/11))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#10](https://github.com/turbot/steampipe-plugin-urlscan/pull/10))

## v0.1.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#5](https://github.com/turbot/steampipe-plugin-urlscan/pull/5))
- Recompiled plugin with Go version 1.17 ([#5](https://github.com/turbot/steampipe-plugin-urlscan/pull/5))

## v0.0.2 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#2](https://github.com/turbot/steampipe-plugin-urlscan/pull/2))


## v0.0.1 [2021-07-01]

_What's new?_

- New tables added
  - [urlscan_app](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_app)
  - [urlscan_asn](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_asn)
  - [urlscan_certificate](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_certificate)
  - [urlscan_console](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_console)
  - [urlscan_cookie](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_cookie)
  - [urlscan_domain](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_domain)
  - [urlscan_global](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_global)
  - [urlscan_ip](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_ip)
  - [urlscan_link](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_link)
  - [urlscan_protocol](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_protocol)
  - [urlscan_request](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_request)
  - [urlscan_resource_type](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_resource_type)
  - [urlscan_search](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_search)
  - [urlscan_server](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_server)
  - [urlscan_task](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_task)
  - [urlscan_tls](https://hub.steampipe.io/plugins/turbot/urlscan/tables/urlscan_tls)
