# Table: urlscan_cookie

List all cookies set by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List cookies

```sql
select
  *
from
  urlscan_cookie
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
```
