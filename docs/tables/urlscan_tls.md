# Table: urlscan_tls

List all TLS protocols used by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List TLS protocols by count

```sql
select
  p.key as protocol,
  p.value as count
from
  urlscan_tls as t,
  jsonb_each(t.protocols) as p
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```
