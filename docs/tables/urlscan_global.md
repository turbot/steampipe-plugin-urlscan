# Table: urlscan_global

List all global variables (properties) in Javascript for the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List global variables

```sql
select
  property,
  type
from
  urlscan_global
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  property
```
