---
title: "Steampipe Table: urlscan_cookie - Query Urlscan Cookies using SQL"
description: "Allows users to query Urlscan Cookies, specifically data related to cookies set during a urlscan.io scan. This provides insights into the cookies associated with a specific URL scan."
---

# Table: urlscan_cookie - Query Urlscan Cookies using SQL

Urlscan.io is a free service that analyzes websites and provides information about their behavior, HTTP requests, and other related details. It offers insights into the structure and security of websites, helping users understand the nature of the sites they visit or analyze. One aspect of this service is the ability to examine cookies set during a scan, providing valuable data about the site's interaction with users' browsers.

## Table Usage Guide

The `urlscan_cookie` table provides insights into the cookies set during a urlscan.io scan. As a security analyst or web developer, you can explore cookie-specific details through this table, including cookie names, values, and associated domains. This table can be utilized to discover information about cookies, such as those related to tracking user behavior, the domains they are associated with, and the potential security implications of these cookies.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List cookies
Explore which cookies are associated with a specific web scan to understand potential tracking or privacy concerns. This can be useful in assessing the security measures of a website or identifying suspicious activity.

```sql+postgres
select
  *
from
  urlscan_cookie
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```

```sql+sqlite
select
  *
from
  urlscan_cookie
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```