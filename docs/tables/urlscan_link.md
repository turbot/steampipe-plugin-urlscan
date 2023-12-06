---
title: "Steampipe Table: urlscan_link - Query Urlscan Links using SQL"
description: "Allows users to query Urlscan Links, particularly the details of each link scanned by Urlscan.io, providing crucial insights into the security status of the links."
---

# Table: urlscan_link - Query Urlscan Links using SQL

Urlscan Links is a resource in Urlscan.io service. It provides detailed information about every link that has been scanned by Urlscan.io. This includes the URL, associated page, and the status of the scan, providing critical insights into the security status of these links.

## Table Usage Guide

The `urlscan_link` table provides insights into the links scanned by Urlscan.io as a security analyst, you can explore the details of each scanned link through this table, including the URL, associated page, and the status of the scan. Utilize it to uncover information about the security status of these links, helping you identify potential security risks and take necessary actions.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List links found in the page
Explore the specific links found within a webpage. This can be particularly useful in identifying potential security risks or understanding the structure and content of a website.

```sql
select
  href,
  text
from
  urlscan_link
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  href
```