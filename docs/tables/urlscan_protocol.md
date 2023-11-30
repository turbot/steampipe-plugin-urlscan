---
title: "Steampipe Table: urlscan_protocol - Query URLScan Protocol using SQL"
description: "Allows users to query Protocols in URLScan, specifically the network protocols used in the HTTP requests made during a scan, providing insights into the security and performance of the network."
---

# Table: urlscan_protocol - Query URLScan Protocol using SQL

URLScan Protocol is a component of URLScan, a web security tool that restricts the types of HTTP requests that Internet Information Services (IIS) will process. It scans HTTP requests to filter out potentially harmful requests, such as those that may exploit a known vulnerability. URLScan Protocol helps to identify and categorize the network protocols used in the HTTP requests made during a scan.

## Table Usage Guide

The `urlscan_protocol` table provides insights into the network protocols used in HTTP requests during a scan in URLScan. As a security analyst, explore protocol-specific details through this table, including the types of protocols used, their frequency, and associated metadata. Utilize it to uncover information about the protocols, such as those frequently used in potentially harmful requests, and the correlation between protocol type and scan results.

## Examples

### List protocols
Explore which protocols are most frequently used in a specific scan to identify potential security risks or unusual network activity. This can help in understanding the network behavior and devise strategies for improved security.

```sql
select
  *
from
  urlscan_protocol
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Protocols by country
Discover the segments that use different protocols, organized by country, to gain insights into internet usage patterns and potential security risks. This information can be useful for identifying areas of concern and optimizing network performance.

```sql
select
  countries ->> 0 as country,
  protocol
from
  urlscan_protocol
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  protocol
```