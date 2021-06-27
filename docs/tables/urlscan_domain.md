# Table: urlscan_domain

List all domains used by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List domains

```sql
select
  *
from
  urlscan_domain
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Domains by country

```sql
select
  countries ->> 0 as country,
  domain
from
  urlscan_domain
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  domain
```
