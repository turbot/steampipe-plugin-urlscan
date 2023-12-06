---
title: "Steampipe Table: urlscan_domain - Query Urlscan Domains using SQL"
description: "Allows users to query Urlscan Domains, particularly the details about domains including the domain name, IP address, and the associated information about the domain's server."
---

# Table: urlscan_domain - Query Urlscan Domains using SQL

Urlscan Domains is a resource within that provides detailed information about domains, including the domain name, IP address, server, and other associated details. It is part of the Urlscan service that offers insights into the behavior and configuration of websites, helping in identifying potential security issues. Urlscan Domains allows you to monitor and analyze the domain-related information effectively.

## Table Usage Guide

The `urlscan_domain` table provides insights into the domains within Urlscan. As a Security Analyst or a DevOps engineer, explore domain-specific details through this table, including the domain name, IP address, and the associated server information. Utilize it to uncover information about domains, such as their server details, the IP address, and other associated metadata.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List domains
Explore which domains are associated with a specific scan in order to understand its reach and impact. This can be useful in identifying potential security threats or analyzing web traffic.

```sql
select
  *
from
  urlscan_domain
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```

### Domains by country
Discover the segments that categorize website domains by their respective countries, useful for understanding geographical distribution and focus of web content.

```sql
select
  countries ->> 0 as country,
  domain
from
  urlscan_domain
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  country,
  domain
```