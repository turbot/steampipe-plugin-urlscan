---
title: "Steampipe Table: urlscan_tls - Query urlscan TLS using SQL"
description: "Allows users to query TLS in urlscan, specifically the detailed information about the TLS certificates of the scanned target, providing insights into the certificate's issuer, subject, validity, and more."
---

# Table: urlscan_tls - Query urlscan TLS using SQL

The Transport Layer Security (TLS) in urlscan provides details about the TLS certificates of a scanned target. It includes information about the certificate's issuer, subject, validity, and other related details. TLS is a security protocol designed to facilitate privacy and data security for communications over the Internet.

## Table Usage Guide

The `urlscan_tls` table provides insights into the TLS certificates of a scanned target within urlscan. As a security analyst, explore detailed information about the certificates through this table, including the certificate's issuer, subject, validity, and more. Utilize it to uncover information about the security of the target, such as the validity of the certificates, the issuer's details, and the associated metadata.

**Important Notes**
- You must specify the `scan` in the `where` clause to query this table.

## Examples

### List TLS protocols by count
Determine the frequency of different TLS protocols for a specific scan to analyze security measures. This can be crucial in identifying potential vulnerabilities and ensuring robust security protocols are in place.

```sql
select
  p.key as protocol,
  p.value as count
from
  urlscan_tls as t,
  jsonb_each(t.protocols) as p
where
  scan = '54c78f69-5294-4a17-8ae0-a71943954e09'
order by
  count desc
```