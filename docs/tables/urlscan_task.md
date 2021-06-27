# Table: urlscan_task

Details of a scanning task.

Note: A `scan` must be provided in all queries to this table.

## Examples

### Get task information

```sql
select
  *
from
  urlscan_task
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
```
