# Table: urlscan_search

Search scan history from urlscan.io to find scans for a domain. The full [urlscan search syntax](https://urlscan.io/docs/search/) is supported in the query field.

Note: A `query` must be provided in all queries to this table.

## Examples

### List all scans for a given domain and it's subdomains

```sql
select
  *
from
  urlscan_search
where
  query = 'domain:steampipe.io'
order by
  task_time desc
```

### List all scans for a given URL (no subdomains)

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:steampipe.io'
order by
  task_time desc
```

### List all scans for an exact page

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:"https://steampipe.io/downloads"'
order by
  task_time desc
```

### Get the last scan for a page

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:"https://steampipe.io/downloads"'
order by
  task_time desc
limit 1
```

### List all scans for a domain in the last 7 days

```sql
select
  *
from
  urlscan_search
where
  query = 'domain:steampipe.io date:now-7d'
order by
  task_time desc
```
