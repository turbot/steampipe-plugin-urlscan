# Table: urlscan_server

List all servers used by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List servers

```sql
select
  *
from
  urlscan_server
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Servers by country

```sql
select
  countries ->> 0 as country,
  server
from
  urlscan_server
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  server
```
