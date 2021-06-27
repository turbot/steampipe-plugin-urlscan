# Table: urlscan_document_type

List all document types returned by requests for the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List document types

```sql
select
  *
from
  urlscan_document_type
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```
