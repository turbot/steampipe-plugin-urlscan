# Table: urlscan_app

List all apps detected in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List detected apps

```sql
select
  *
from
  urlscan_app
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
```
