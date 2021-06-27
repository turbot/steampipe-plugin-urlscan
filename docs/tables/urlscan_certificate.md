# Table: urlscan_certificate

List all certificates found during a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List certificates found in the scan

```sql
select
  subject_name,
  issuer,
  valid_to,
  (valid_to::date - current_date) as days_until_expiration
from
  urlscan_certificate
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  valid_to
```
