---
title: "Steampipe Table: urlscan_server - Query URLscan Servers using SQL"
description: "Allows users to query URLscan Servers, specifically providing insights into server details and potential security vulnerabilities."
---

# Table: urlscan_server - Query URLscan Servers using SQL

URLscan is a service that allows users to analyze and monitor the security of their web servers. It provides detailed information about the server, including IP addresses, locations, and security vulnerabilities. With URLscan, users can easily identify potential security threats and take appropriate actions to mitigate them.

## Table Usage Guide

The `urlscan_server` table provides detailed insights into the servers scanned by URLscan. As a security analyst, explore server-specific details through this table, including IP addresses, locations, and security vulnerabilities. Utilize this table to uncover potential security threats and take appropriate actions to mitigate them.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List servers
Gain insights into the servers associated with a particular scan to understand its activity and usage. This can be useful for monitoring server performance, identifying potential issues, or tracking the impact of specific scans.

```sql
select
  *
from
  urlscan_server
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Servers by country
Explore which servers are associated with different countries. This is particularly useful for understanding the geographical distribution of your servers and identifying potential regional issues.

```sql
select
  countries ->> 0 as country,
  server
from
  urlscan_server
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  server
```