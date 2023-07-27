![image](https://hub.steampipe.io/images/plugins/turbot/urlscan-social-graphic.png)

# urlscan.io Plugin for Steampipe

Use SQL to query URL scanning results including requests, cookies, headers and more from urlscan.io.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/urlscan)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/urlscan/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-urlscan/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install urlscan
```

Run a query:

```sql
select
  id,
  page_url,
  task_time
from
  urlscan_search
where
  query = 'steampipe.io'
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-urlscan.git
cd steampipe-plugin-urlscan
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/urlscan.spc
```

Try it!

```
steampipe query
> .inspect urlscan
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-urlscan/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [urlscan.io Plugin](https://github.com/turbot/steampipe-plugin-urlscan/labels/help%20wanted)
