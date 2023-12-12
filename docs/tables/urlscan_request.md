---
title: "Steampipe Table: urlscan_request - Query Urlscan Requests using SQL"
description: "Allows users to query Urlscan Requests, which provides valuable information about the HTTP request made during the scanning process."
---

# Table: urlscan_request - Query Urlscan Requests using SQL

Urlscan is a tool used to analyze and scan websites for potential security threats. It checks various aspects of a webpage, including HTTP requests made during the scanning process. These requests are crucial for understanding the nature of the webpage and its potential security implications.

## Table Usage Guide

The `urlscan_request` table provides insights into the HTTP requests made during the scanning process by Urlscan. As a Security Analyst, you can use this table to uncover detailed information about these requests, such as the method, URL, and headers. This can assist in identifying potential security threats and understanding the behavior of the scanned webpage.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List requests
Analyze the sequence of requests made during a specific web scan to understand the progression and interactions within that session. This could be useful in identifying potential security threats or anomalies in the browsing session.

```sql+postgres
select
  *
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  timestamp;
```

```sql+sqlite
select
  *
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  timestamp;
```

### Post requests made while loading the page
Explore which URLs were accessed during a specific page load, particularly focusing on POST requests. This is beneficial to understand the data being transmitted during the loading process, which can aid in troubleshooting or security analysis.

```sql+postgres
select
  url,
  type,
  response_status,
  post_data
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and has_post_data;
```

```sql+sqlite
select
  url,
  type,
  response_status,
  post_data
from
  urlscan_request
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and has_post_data = 1;
```

### 5 largest requests in the page
Determine the areas in which your webpage is making the largest requests to help focus optimization efforts and improve loading times.

```sql+postgres
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
limit 5;
```

```sql+sqlite
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
limit 5;
```

### Types of documents requested in page
Assess the elements within a specific webpage scan to understand the frequency and total size of different document types requested. This can be useful in identifying potential areas of optimization to improve webpage loading speeds.

```sql+postgres
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
  total_size_bytes desc;
```

```sql+sqlite
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
  total_size_bytes desc;
```