# Table: urlscan_link

List all links found in the page for a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List links found in the page

```sql
select
  href,
  text
from
  urlscan_link
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  href
```
