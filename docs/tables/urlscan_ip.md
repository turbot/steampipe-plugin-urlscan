# Table: urlscan_ip

List all IPs used by requests in the page of a given scan.

Note: A `scan` must be provided in all queries to this table.

## Examples

### List IP addresses

```sql
select
  *
from
  urlscan_ip
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### IPs by country

```sql
select
  countries ->> 0 as country,
  ip
from
  urlscan_ip
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  ip
```

### IPs with Geolocation

```sql
select
  geolocation ->> 'country_name' as country,
  geolocation ->> 'region' as region,
  geolocation ->> 'city' as city,
  ip
from
  urlscan_ip
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  region,
  city,
  ip
```
