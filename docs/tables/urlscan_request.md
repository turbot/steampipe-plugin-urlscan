# Table: urlscan_request

List all requests made to build the requested page, including JS, CSS, etc.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List requests

```sql
select
  *
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  timestamp
```

### Post requests made while loading the page

```sql
select
  url,
  type,
  response_status,
  post_data
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and has_post_data
```

### 5 largest requests in the page

```sql
select
  response_size,
  document_type,
  url
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  response_size desc
limit 5
```

### Types of documents requested in page

```sql
select
  document_type,
  count(*) as count,
  sum(response_size) as total_size_bytes
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
group by
  document_type
order by
  total_size_bytes desc
```
