---
title: "Steampipe Table: urlscan_ip - Query Urlscan IP Addresses using SQL"
description: "Allows users to query IP Addresses in Urlscan, specifically to retrieve information about the IP addresses associated with a scan, providing insights into the IP addresses that a website connects to."
---

# Table: urlscan_ip - Query Urlscan IP Addresses using SQL

Urlscan is a free and open API and web interface for scanning and analyzing web sites. It checks for issues with security, privacy, and compliance, and provides detailed information about the website's connections and the resources loaded. An IP Address in Urlscan represents a unique numerical label assigned to each device participating in a computer network that uses the Internet Protocol for communication.

## Table Usage Guide

The `urlscan_ip` table provides insights into IP addresses associated with a scan in Urlscan. As a security analyst, explore details about these IP addresses through this table, including their geographical location, associated domain, and ASN details. Utilize it to uncover information about the IP addresses a website connects to, helping to identify potential security risks or malicious activity.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List IP addresses
Explore the IP addresses associated with a particular scan to understand its reach and impact. This can be used to determine potential security threats or anomalies in the network traffic.

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
Explore which IP addresses are associated with specific countries to enhance your understanding of your network's geographical distribution. This could be crucial for detecting unusual activity or potential security threats.

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
Explore the geographical locations associated with specific IP addresses. This is useful for identifying patterns or anomalies in network traffic, potentially highlighting security threats or operational issues.

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