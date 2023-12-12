---
title: "Steampipe Table: urlscan_asn - Query Urlscan ASN using SQL"
description: "Allows users to query Autonomous System Numbers (ASNs) in Urlscan, specifically providing insights into the network routing prefixes and the networks they belong to."
---

# Table: urlscan_asn - Query Urlscan ASN using SQL

An Autonomous System Number (ASN) is a unique identifier assigned to an Autonomous System (AS). ASNs are used by the Border Gateway Protocol (BGP) for routing traffic across the internet. In Urlscan, ASNs provide insights into the network routing prefixes and the networks they belong to.

## Table Usage Guide

The `urlscan_asn` table provides insights into Autonomous System Numbers (ASNs) within Urlscan. As a network engineer, explore ASN-specific details through this table, including the network routing prefixes and the networks they belong to. Utilize it to uncover information about ASNs, such as their associated networks, and the verification of routing prefixes.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List ASNs
Explore the Autonomous System Numbers (ASNs) related to a specific web scan. This is useful for identifying the networks involved in the scanned web activity, contributing to enhanced cybersecurity measures.

```sql+postgres
select
  *
from
  urlscan_asn
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```

```sql+sqlite
select
  *
from
  urlscan_asn
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09';
```