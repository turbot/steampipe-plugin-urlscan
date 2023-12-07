---
title: "Steampipe Table: urlscan_app - Query UrlScan Apps using SQL"
description: "Allows users to query UrlScan Apps, specifically providing insights into app data including its associated tags, related URLs, and the app's overall popularity."
---

# Table: urlscan_app - Query UrlScan Apps using SQL

UrlScan is a service within that allows users to scan and analyze web applications for potential security vulnerabilities. It provides a comprehensive view of the application's security posture, including details about associated tags, related URLs, and the app's overall popularity. UrlScan helps users stay informed about the health and performance of their web applications and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `urlscan_app` table provides insights into web applications within UrlScan. As a security analyst, explore app-specific details through this table, including associated tags, related URLs, and the app's overall popularity. Utilize it to uncover information about apps, such as those with potential security vulnerabilities, the relationships between apps and URLs, and the verification of app popularity.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List detected apps
Explore which applications have been detected by analyzing the results of a specific scan. This could be useful in identifying potential security threats or understanding the software landscape of a system.

```sql+postgres
select
  *
from
  urlscan_app
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```

```sql+sqlite
select
  *
from
  urlscan_app
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```