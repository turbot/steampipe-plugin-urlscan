# Table: urlscan_protocol

List all protocols used by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List protocols

```sql
select
  *
from
  urlscan_protocol
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Protocols by country

```sql
select
  countries ->> 0 as country,
  protocol
from
  urlscan_protocol
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  protocol
```
