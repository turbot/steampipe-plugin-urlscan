# Table: urlscan_metadata

Metadata details of a scanning.

Note: A `scan` must be provided in all queries to this table.

## Examples

### Get metadata details for a scanning

```sql
select
  *
from
  urlscan_metadata
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```

### List asn metadata details

```sql
select
  scan,
  data -> 'country' as country,
  data -> 'date' as date,
  data -> 'ip' as ip,
  data -> 'name' as name,
  data -> 'register' as register,
  data -> 'route' as route
from
  urlscan_metadata,
  jsonb_array_elements(asn_data) as data
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```
