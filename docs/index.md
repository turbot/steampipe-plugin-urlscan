---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/urlscan.svg"
brand_color: "#D2624D"
display_name: "urlscan.io"
short_name: "urlscan"
description: "Steampipe plugin to query URL scanning results including requests cookies, headers and more from urlscan.io."
og_description: "Query urlscan.io with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/urlscan-social-graphic.png"
---

# urlscan.io + Steampipe

[urlscan.io](https://urlscan.io) is a website scanner to gather request, cookie, link, headers and other information from a URL.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Search website scans:

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

```
+--------------------------------------+---------------------------+---------------------+
| id                                   | page_url                  | task_time           |
+--------------------------------------+---------------------------+---------------------+
| a20397de-1f23-4821-ae2a-87ca612ebcff | https://steampipe.io/     | 2021-05-06 01:39:14 |
| 698e6270-cd8f-4090-a3ae-23bb5d3ee0a8 | https://steampipe.io/     | 2020-12-15 19:44:10 |
| d01fa12c-f30f-48b8-af7e-5539f0d1bc3c | https://hub.steampipe.io/ | 2020-12-04 11:44:02 |
+--------------------------------------+---------------------------+---------------------+
```

Query URL information from a scan:

```sql
select
  subject_name,
  issuer,
  valid_to
from
  urlscan_certificate
where
  -- vercel.com
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
```

```
+-------------------------------+--------------------------------------------+---------------------+
| subject_name                  | issuer                                     | valid_to            |
+-------------------------------+--------------------------------------------+---------------------+
| *.vercel.com                  | R3                                         | 2021-06-15 03:05:04 |
| vercel-insights.com           | Amazon                                     | 2021-12-22 23:59:59 |
| *.google.com                  | GTS CA 1O1                                 | 2021-07-06 10:11:14 |
| ...                           | ...                                        | ...                 |
+-------------------------------+--------------------------------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/urlscan/tables)**

## Get started

### Install

Download and install the latest urlscan.io plugin:

```bash
steampipe plugin install urlscan
```

### Credentials

| Item        | Description                                                                             |
| :---------- | :-------------------------------------------------------------------------------------- |
| Credentials | urlscan.io requires a [free API key](https://urlscan.io/user/signup/) for all requests. |
| Radius      | Each connection represents a single urlscan.io account.                                 |

### Configuration

Installing the latest urlscan plugin will create a config file (`~/.steampipe/config/urlscan.spc`) with a single connection named `urlscan`:

```hcl
connection "urlscan" {
  plugin  = "urlscan"
  api_key = "e89d13c7-43a0-4f48-a829-590b4091bac6"
}
```

- `api_key` - Your urlscan.io API key.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-urlscan
- Community: [Slack Channel](https://steampipe.io/community/join)
