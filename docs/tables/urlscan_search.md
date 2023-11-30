---
title: "Steampipe Table: urlscan_search - Query Urlscan Search Results using SQL"
description: "Allows users to query Urlscan Search Results, specifically providing insights into the detailed analysis of websites and webpages."
---

# Table: urlscan_search - Query Urlscan Search Results using SQL

Urlscan is a tool that provides detailed analysis of websites and webpages. It scans and checks for malicious content, providing a comprehensive view of how a website is connected to other domains and subdomains. It is particularly useful for identifying and understanding threats, malicious scripts, and other suspicious activities associated with the website.

## Table Usage Guide

The `urlscan_search` table provides insights into the detailed analysis of websites and webpages using Urlscan. As a security analyst, explore specific details through this table, including website connections, malicious content, and associated metadata. Utilize it to uncover information about potential threats, malicious scripts, and other suspicious activities associated with the website.

## Examples

### List all scans for a given domain and it's subdomains
Explore all scans associated with a specific domain and its subdomains, including those most recently conducted. This is useful for tracking and analyzing the security and performance history of a website.

```sql
select
  *
from
  urlscan_search
where
  query = 'domain:steampipe.io'
order by
  task_time desc
```

### List all scans for a given URL (no subdomains)
Uncover the details of all scans conducted for a specific website, ordered by the time they were performed. This can help in analyzing the security and performance history of the website.

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:steampipe.io'
order by
  task_time desc
```

### List all scans for an exact page
Explore which scans have been conducted for a specific webpage. This is useful for tracking the frequency and timing of scans, helping to monitor site security and performance over time.

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:"https://steampipe.io/downloads"'
order by
  task_time desc
```

### Get the last scan for a page
Analyze the settings to understand the most recent scan performed on a specific webpage. This is useful for tracking changes over time or identifying potential issues or vulnerabilities that may have been introduced in the latest update.

```sql
select
  *
from
  urlscan_search
where
  query = 'page.url:"https://steampipe.io/downloads"'
order by
  task_time desc
limit 1
```

### List all scans for a domain in the last 7 days
Explore the recent security scans performed on a specific domain over the past week. This can help in identifying potential threats and maintaining the security integrity of the domain.

```sql
select
  *
from
  urlscan_search
where
  query = 'domain:steampipe.io date:now-7d'
order by
  task_time desc
```