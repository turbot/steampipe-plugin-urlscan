---
title: "Steampipe Table: urlscan_console - Query Urlscan Console using SQL"
description: "Allows users to query Urlscan Consoles, specifically providing insights into scan results, associated metadata, and related information."
---

# Table: urlscan_console - Query Urlscan Console using SQL

Urlscan Console is a part of Oracle Cloud Infrastructure's security services. It allows users to perform security scans and provides detailed analysis of the scan results. The console aids in identifying security threats, vulnerabilities, and helps in maintaining the overall security posture of your resources.

## Table Usage Guide

The `urlscan_console` table provides insights into the scan results within Urlscan Console. As a security analyst, explore scan-specific details through this table, including vulnerabilities, associated metadata, and related information. Utilize it to uncover information about potential threats and vulnerabilities, and to verify the security posture of your resources.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List console messages in order
Analyze the sequence of console messages to gain insights into the progression of events for a specific scan. This can be useful for troubleshooting or understanding the behavior of a system.

```sql+postgres
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  rank;
```

```sql+sqlite
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  rank;
```

### List warnings and errors from the console
Explore the instances where warnings and errors have been logged in a specific system scan. This query is useful in identifying potential issues and prioritizing them based on their severity for swift remediation.

```sql+postgres
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and level in ('warning', 'error')
order by
  rank;
```

```sql+sqlite
select
  *
from
  urlscan_console
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
  and level in ('warning', 'error')
order by
  rank;
```