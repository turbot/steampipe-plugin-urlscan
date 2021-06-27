# Table: urlscan_console

List all console messages generated during the page load for a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List console messages in order

```sql
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  rank
```

### List warnings and errors from the console

```sql
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and level in ('warning', 'error')
order by
  rank
```
